package storage

// API 存放接口信息的表
type API struct {
	Id       int64  `json:"id"`
	Method   string `gorm:"size:128;not null;default:GET" json:"method"`
	Path     string `gorm:"size:128;not null;uniqueIndex" json:"path"`
	Header   string `gorm:"type:text" json:"header"`
	Response string `gorm:"type:text" json:"response"`
}

// User 用户表
type User struct {
	Id         int64  `json:"id"`
	UserName   string `gorm:"size:128;not null;column:username;unique" json:"username"`
	PassWord   string `gorm:"not null;column:password" json:"password"`
	EMail      string `gorm:"size:256;column:e-mail" json:"e-mail"`
	CommonPath string `gorm:"size:256" json:"common_path"`
}
