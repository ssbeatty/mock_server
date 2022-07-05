package web

const (
	HttpStatusOk               = 200
	HttpStatusBadRequest       = 400
	HttpStatusNotFound         = 404
	HttpResponseSuccess        = "success"
	HTTPResponseRouterNotFound = "router not found"
)

// Response 通用http请求response
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

func generateResponsePayload(code int, msg string, data ...interface{}) Response {
	if len(data) > 0 {
		return Response{code, msg, data[0]}
	}
	return Response{code, msg, nil}
}

// RegisterPayload 注册payload
type RegisterPayload struct {
	UserName string `json:"username" binding:"required"`
	PassWord string `json:"password" binding:"required"`
	Email    string `json:"e-mail"`
}

// GetRouterParam 获取API详情的参数
type GetRouterParam struct {
	Id int64 `uri:"id" binding:"required"`
}

// PostRouterForm 新增API
type PostRouterForm struct {
	Method   string `json:"method" binding:"required"`
	Path     string `json:"path" binding:"required"`
	Header   string `json:"header"`
	Response string `json:"response"`
}

// PutRouterForm 更新API
type PutRouterForm struct {
	Id       int64  `uri:"id" binding:"required"`
	Method   string `json:"method" binding:"required"`
	Path     string `json:"path" binding:"required"`
	Header   string `json:"header"`
	Response string `json:"response"`
}

// DeleteRouterParam 删除API
type DeleteRouterParam struct {
	Id int64 `uri:"id" binding:"required"`
}
