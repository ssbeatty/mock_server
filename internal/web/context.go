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

func (c *Context) ResponseError(msg string) {
	c.JSON(http.StatusOK, msg)
}

func (c *Context) ResponseOk(data interface{}) {
	c.JSON(http.StatusOK, data)
}
