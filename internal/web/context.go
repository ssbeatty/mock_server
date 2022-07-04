package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerFunc func(c *Context)

func Handle(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &Context{
			Context: c,
		}
		h(ctx)
	}
}

type Context struct {
	*gin.Context
}

func (c *Context) ResponseError(msg string, code int) {
	d := generateResponsePayload(code, msg)
	c.JSON(http.StatusOK, d)
}

func (c *Context) ResponseOk(data ...interface{}) {
	d := generateResponsePayload(HttpStatusOk, HttpResponseSuccess, data...)
	c.JSON(http.StatusOK, d)
}
