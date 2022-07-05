package storage

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/logger"
	"testing"
)

var service *Service

func TestMain(m *testing.M) {
	service = NewService(&DB{
		Driver:   "postgres",
		Dsn:      "127.0.0.1:5432",
		DbName:   "mock_server",
		UserName: "odoo",
		PassWord: "odoo",
	})
	service.Open()
	service.db.Logger = logger.Default.LogMode(logger.Info)

	m.Run()
}

func TestService_SaveUser(t *testing.T) {
	err := service.SaveUser("demo", "demo", "")
	if err != nil {
		t.Error(err)
	}
	assert.Nil(t, err)
}
