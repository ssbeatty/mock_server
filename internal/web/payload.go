package web

const (
	HttpStatusOk         = 200
	HttpStatusBadRequest = 400
	HttpResponseSuccess  = "success"
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
