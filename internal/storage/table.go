package storage

import "gorm.io/gorm"

type API struct {
	gorm.Model

	Id       int    `json:"id"`
	Path     string `gorm:"size:128;not null" json:"path"`
	Header   string `gorm:"type:text" json:"header"`
	Response string `gorm:"type:text" json:"response"`
}
