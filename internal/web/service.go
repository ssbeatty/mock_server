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

	db     *storage.Service
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

	s.Logger.Info("web服务加载成功.")
	return s.Run(s.Addr)
}

func (s *Service) initRouters() {
	s.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
