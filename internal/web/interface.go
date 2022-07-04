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
	GetAllRouters() ([]storage.API, error)
	GetUserByName(username string) (*storage.User, error)
	SaveUser(username, password, email string) error
}
