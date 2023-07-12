package controller

import "github.com/gin-gonic/gin"

type IOrderController interface {
	CreateOrder(*gin.Context)
}

type orderController struct {
}

func NewRouter() IOrderController {
	return &orderController{}
}

// ----------------------------------

func (*orderController) CreateOrder(ctx *gin.Context) {

}
