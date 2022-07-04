package web

func (s *Service) GetRouters(c *Context) {
	hosts, err := s.db.GetAllRouters()
	if err != nil {
		s.Logger.Errorf("获取所有接口信息出错: %v", err)
		c.ResponseError(err.Error())
		return
	}
	c.ResponseOk(hosts)
}
