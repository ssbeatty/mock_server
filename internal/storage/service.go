package storage

import (
	"errors"
	"fmt"
	"github.com/Pacific73/gorm-cache/cache"
	"github.com/Pacific73/gorm-cache/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mock_server/pkg/log"
	"strings"
)

type DB struct {
	Driver   string `yaml:"driver"`
	UserName string `yaml:"username"`
	PassWord string `yaml:"password"`
	Dsn      string `yaml:"dsn"`
	DbName   string `yaml:"dbname"`
}

type Service struct {
	db     *gorm.DB
	logger log.Logger
	conf   *DB
}

func NewService(conf *DB) *Service {

	return &Service{
		conf: conf,
	}
}

func (s *Service) Open() error {
	var d *gorm.DB
	var err error
	var dataSource string

	if s.conf.Driver == "mysql" {
		dataSource = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", s.conf.UserName, s.conf.PassWord, s.conf.Dsn, s.conf.DbName)
		d, err = gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	} else if s.conf.Driver == "postgres" {
		dsnArgs := strings.Split(s.conf.Dsn, ":")
		if len(dsnArgs) < 2 {
			return errors.New("dsn parse error")
		}
		dataSource = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			dsnArgs[0], s.conf.UserName, s.conf.PassWord, s.conf.DbName, dsnArgs[1],
		)
		d, err = gorm.Open(postgres.Open(dataSource), &gorm.Config{})
	} else {
		d, err = nil, errors.New("driver is not support")
	}
	if err != nil {
		return err
	}
	if allCache, err := cache.NewGorm2Cache(&config.CacheConfig{
		CacheLevel:           config.CacheLevelAll,
		CacheStorage:         config.CacheStorageMemory,
		InvalidateWhenUpdate: true,
		CacheTTL:             24 * 60 * 60 * 1000,
		CacheMaxItemCnt:      5000,
		CacheSize:            5000,
		DebugMode:            false,
	}); err == nil {
		err := d.Use(allCache)
		if err != nil {
			return err
		}
	}

	s.db = d
	if err = s.db.AutoMigrate(new(API)); err != nil {
		return err
	}

	return nil
}
