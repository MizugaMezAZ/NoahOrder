package router

import "github.com/gin-gonic/gin"

type router struct {
}

func NewRouter() *router {
	return &router{}
}

func (r *router) SetupRoute(e *gin.Engine) {
	e.Use()
}
