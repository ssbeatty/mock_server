package web

// 用户相关接口

// Register 注册
func (s *Service) Register(c *Context) {
	var form RegisterPayload
	err := c.ShouldBindJSON(&form)
	if err != nil {
		c.ResponseError(err.Error(), HttpStatusBadRequest)
	} else {
		err := s.db.SaveUser(form.UserName, form.PassWord, form.Email)
		if err != nil {
			c.ResponseError(err.Error(), HttpStatusBadRequest)
			return
		}
		c.ResponseOk()
	}
}

// GetRouters 获取所有的接口
func (s *Service) GetRouters(c *Context) {
	hosts, err := s.db.GetAllRouters()
	if err != nil {
		s.Logger.Errorf("获取所有接口信息出错: %v", err)
		c.ResponseError(err.Error(), HttpStatusBadRequest)
		return
	}
	c.ResponseOk(hosts)
}
