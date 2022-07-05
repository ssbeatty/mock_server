package web

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"mock_server/config"
	"mock_server/internal/storage"
	"mock_server/pkg/log"
	"net/http"
	"time"
)

const (
	identityKey     = "id"
	JWTTokenTimeout = time.Hour
	JWTMaxRefresh   = 14 * 24 * time.Hour
)

type Service struct {
	*gin.Engine

	secretPath string
	db         IStorage
	Addr       string
	Logger     Logger
}

type login struct {
	UserName string `form:"username" json:"username" binding:"required"`
	PassWord string `form:"password" json:"password" binding:"required"`
}

func NewService(app *config.App, db *storage.Service) *Service {
	gin.SetMode(gin.ReleaseMode)
	logger := log.NewLogrusAdapt(logrus.StandardLogger()).WithField("service", "web")

	service := &Service{
		Addr:       app.Addr,
		secretPath: app.SecretKeyPath,
		Logger:     logger,
		Engine:     gin.Default(),
		db:         db,
	}

	return service
}

func (s *Service) Serve() error {
	s.initRouters()

	s.Logger.Infof("web服务加载成功. Listen: %s", s.Addr)
	return s.Run(s.Addr)
}

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID":   claims[identityKey],
		"userName": user.(*storage.User).UserName,
		"text":     "Hello World.",
	})
}

func (s *Service) initRouters() {
	s.Use(CORS).Use(exportHeaders)

	s.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	data, err := config.ParseSecret(s.secretPath)
	if err != nil {
		s.Logger.Fatal("获取secret key失败!")
	}

	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "Mock Server",
		Key:         data,
		Timeout:     JWTTokenTimeout,
		MaxRefresh:  JWTMaxRefresh,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*storage.User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			user, err := s.db.GetUserByName(claims[identityKey].(string))
			if err != nil {
				return nil
			}
			return user
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.UserName
			password := loginVals.PassWord

			user, err := s.db.GetUserByName(userID)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			err = bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(password)) //验证（对比）
			if err != nil {
				return nil, err
			} else {
				return user, nil
			}
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		SendCookie:    true,
	})

	if err != nil {
		s.Logger.Fatal("JWT Error:" + err.Error())
	}

	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		s.Logger.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	// 用户相关接口
	s.POST("/login", authMiddleware.LoginHandler)
	s.POST("/register", Handle(s.Register))
	auth := s.Group("/auth").Use(authMiddleware.MiddlewareFunc())
	{
		// Refresh time can be longer than token timeout
		auth.POST("/refresh_token", authMiddleware.RefreshHandler)
		auth.POST("/logout", authMiddleware.LogoutHandler)
		auth.GET("/hello", helloHandler)
	}

	// 管理相关接口
	apiV1 := s.Group("/admin/api/v1").
		Use(authMiddleware.MiddlewareFunc())
	{
		apiV1.GET("/routers", Handle(s.GetRouters))
		apiV1.GET("/router/:id", Handle(s.GetRouter))
		apiV1.POST("/router", Handle(s.CreateRouter))
		apiV1.PUT("/router", Handle(s.UpdateRouter))
		apiV1.DELETE("/router/:id", Handle(s.DeleteRouter))
	}

	s.Any("/mock/*path", Handle(s.RouteingMapIndex))
}
