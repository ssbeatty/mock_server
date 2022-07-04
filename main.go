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

	addr := viper.GetString("addr")

	db := &storage.DB{}
	err := viper.UnmarshalKey("db", db)
	if err != nil {
		log.Fatal(err)
	}

	dbSrv := storage.NewService(db)
	err = dbSrv.Open()
	if err != nil {
		log.Fatal(err)
	}

	service := web.NewService(addr, dbSrv)

	if err := service.Serve(); err != nil {
		log.Fatal(err)
	}
}
