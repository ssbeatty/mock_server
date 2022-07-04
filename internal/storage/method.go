package storage

// GetAllRouters 获取所有的接口列表
func (s *Service) GetAllRouters() ([]API, error) {
	var routers []API
	err := s.db.Find(&routers).Error
	if err != nil {
		return nil, err
	}

	return routers, nil
}
