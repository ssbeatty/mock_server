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

// GetRouterById 通过ID获取路由信息
func (s *Service) GetRouterById(id int64) (*API, error) {
	var router API
	err := s.db.First(&router, id).Error
	if err != nil {
		return nil, err
	}

	return &router, nil
}

// GetRouterByPath 通过Path获取路由信息
func (s *Service) GetRouterByPath(path string) (*API, error) {
	var router API
	err := s.db.Where("path = ?", path).First(&router).Error
	if err != nil {
		return nil, err
	}

	return &router, nil
}

// CreateRouter 新增路由信息
func (s *Service) CreateRouter(method, path, header, response string) (*API, error) {
	router := API{
		Method:   method,
		Path:     path,
		Header:   header,
		Response: response,
	}
	err := s.db.Create(&router).Error
	if err != nil {
		return nil, err
	}

	return &router, nil
}

// UpdateRouter 修改路由信息
func (s *Service) UpdateRouter(id int64, method, path, header, response string) (*API, error) {
	router := API{
		Id:       id,
		Method:   method,
		Path:     path,
		Header:   header,
		Response: response,
	}
	err := s.db.Updates(&router).Error
	if err != nil {
		return nil, err
	}

	return &router, nil
}

// DeleteRouter 删除路由
func (s *Service) DeleteRouter(id int64) error {
	return s.db.Delete(&API{}, id).Error
}

// 用户相关接口

// GetUserByName 通过用户名获取用户
func (s *Service) GetUserByName(username string) (*User, error) {
	var user User
	err := s.db.Where("UserName = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// SaveUser 新增用户
// db.Save 会更新0值字段
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
		UserName:   username,
		PassWord:   string(hash),
		EMail:      email,
		CommonPath: username,
	}
	err = s.db.Save(user).Error
	if err != nil {
		return err
	}

	return nil
}

// UpdateUser 更新用户信息
func (s *Service) UpdateUser(user *User) error {
	return s.db.Updates(user).Error
}
