package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"mock_server/pkg/log"
)

type Service struct {
	*gin.Engine
	Addr   string
	Port   int
	Logger Logger
}

func NewService(Addr string, Port int) *Service {
	service := &Service{
		Addr:   fmt.Sprintf("%s:%d", Addr, Port),
		Logger: log.NewLogrusAdapt(logrus.StandardLogger()),
	}

	return service
}
