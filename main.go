package main

import (
	"github.com/spf13/viper"
	"log"
	"mock_server/config"
	"mock_server/internal/storage"
	"mock_server/internal/web"
)

func main() {
	if err := config.NewConfig(); err != nil {
		log.Fatal(err)
	}

	app := &config.App{}
	if err := viper.UnmarshalKey("app", app); err != nil {
		log.Fatal(err)
	}

	db := &storage.DB{}
	if err := viper.UnmarshalKey("db", db); err != nil {
		log.Fatal(err)
	}

	dbSrv := storage.NewService(db)
	err := dbSrv.Open()
	if err != nil {
		log.Fatal(err)
	}

	service := web.NewService(app, dbSrv)

	if err := service.Serve(); err != nil {
		log.Fatal(err)
	}
}
