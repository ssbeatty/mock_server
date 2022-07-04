package web

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"mock_server/internal/storage"
	"mock_server/pkg/log"
	"net/http"
)

type Service struct {
	*gin.Engine

	db     IStorage
	Addr   string
	Logger Logger
}

func NewService(Addr string, db *storage.Service) *Service {
	gin.SetMode(gin.ReleaseMode)
	logger := log.NewLogrusAdapt(logrus.StandardLogger()).WithField("service", "web")

	service := &Service{
		Addr:   Addr,
		Logger: logger,
		Engine: gin.Default(),
		db:     db,
	}

	return service
}

func (s *Service) Serve() error {
	s.initRouters()

	s.Logger.Infof("web服务加载成功. Listen: %s", s.Addr)
	return s.Run(s.Addr)
}

func (s *Service) initRouters() {
	s.Use(CORS).Use(exportHeaders)

	s.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	// admin apis
	apiV1 := s.Group("/admin/api/v1")
	{
		apiV1.GET("/routers", Handle(s.GetRouters))
	}
}
