package storage

// API 存放接口信息的表
type API struct {
	Id       int    `json:"id"`
	Path     string `gorm:"size:128;not null" json:"path"`
	Header   string `gorm:"type:text" json:"header"`
	Response string `gorm:"type:text" json:"response"`
}
