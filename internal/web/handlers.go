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

// 接口管理CURD

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

// GetRouter 获取单个接口
func (s *Service) GetRouter(c *Context) {
	var param GetRouterParam
	err := c.ShouldBindUri(&param)
	if err != nil {
		c.ResponseError(err.Error(), HttpStatusBadRequest)
	} else {
		host, err := s.db.GetRouterById(param.Id)
		if err != nil {
			s.Logger.Errorf("获取接口ID: %d 信息出错: %v", param.Id, err)
			c.ResponseError(err.Error(), HttpStatusBadRequest)
			return
		}
		c.ResponseOk(host)
	}
}

// CreateRouter 创建接口
func (s *Service) CreateRouter(c *Context) {
	var param PostRouterForm
	err := c.ShouldBindJSON(&param)
	if err != nil {
		c.ResponseError(err.Error(), HttpStatusBadRequest)
	} else {
		router, err := s.db.CreateRouter(param.Method, param.Path, param.Header, param.Response)
		if err != nil {
			s.Logger.Errorf("创建接口出错: %v", err)
			c.ResponseError(err.Error(), HttpStatusBadRequest)
			return
		}
		c.ResponseOk(router)
	}
}

// UpdateRouter 修改接口
func (s *Service) UpdateRouter(c *Context) {
	var param PutRouterForm
	err := c.ShouldBind(&param)
	if err != nil {
		c.ResponseError(err.Error(), HttpStatusBadRequest)
	} else {
		router, err := s.db.UpdateRouter(param.Id, param.Method, param.Path, param.Header, param.Response)
		if err != nil {
			s.Logger.Errorf("修改接口出错: %v", err)
			c.ResponseError(err.Error(), HttpStatusBadRequest)
			return
		}
		c.ResponseOk(router)
	}
}

// DeleteRouter 删除接口
func (s *Service) DeleteRouter(c *Context) {
	var param DeleteRouterParam
	err := c.ShouldBindUri(&param)
	if err != nil {
		c.ResponseError(err.Error(), HttpStatusBadRequest)
	} else {
		err := s.db.DeleteRouter(param.Id)
		if err != nil {
			s.Logger.Errorf("删除接口ID: %d 出错: %v", param.Id, err)
			c.ResponseError(err.Error(), HttpStatusBadRequest)
			return
		}
		c.ResponseOk()
	}
}

// RouteingMapIndex 所有的接口的map路由handlers
func (s *Service) RouteingMapIndex(c *Context) {
	path := c.Param("path")

	router, err := s.db.GetRouterByPath(path)
	if err != nil {
		s.Logger.Errorf("未注册的接口path: %s", path)
		c.ResponseError(HTTPResponseRouterNotFound, HttpStatusNotFound)
		return
	}

	c.Writer.WriteString(router.Response)
}
