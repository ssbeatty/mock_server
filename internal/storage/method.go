package storage

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// API相关接口

// GetAllRouters 获取所有的接口列表
func (s *Service) GetAllRouters() ([]API, error) {
	var routers []API
	err := s.db.Find(&routers).Error
	if err != nil {
		return nil, err
	}

	return routers, nil
}

// 用户相关接口

func (s *Service) GetUserByName(username string) (*User, error) {
	var user User
	err := s.db.Where("UserName = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// SaveUser 保存或新增用户
func (s *Service) SaveUser(username, password, email string) error {
	u, err := s.GetUserByName(username)
	if err == nil && u != nil {
		return errors.New("用户已存在")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		return err
	}
	user := &User{
		UserName: username,
		PassWord: string(hash),
		EMail:    email,
	}
	err = s.db.Save(user).Error
	if err != nil {
		return err
	}

	return nil
}
