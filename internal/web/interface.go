package web

import (
	"mock_server/internal/storage"
	"mock_server/pkg/log"
)

type Logger interface {
	WithField(key string, value interface{}) log.Logger
	WithFields(fields log.Fields) log.Logger
	Trace(args ...interface{})
	Tracef(format string, args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

type IStorage interface {
	GetUserByName(username string) (*storage.User, error)
	SaveUser(username, password, email string) error
	UpdateUser(user *storage.User) error

	GetAllRouters() ([]storage.API, error)
	GetRouterById(id int64) (*storage.API, error)
	GetRouterByPath(path string) (*storage.API, error)
	CreateRouter(method, path, header, response string) (*storage.API, error)
	UpdateRouter(id int64, method, path, header, response string) (*storage.API, error)
	DeleteRouter(id int64) error
}
