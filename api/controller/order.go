package controller

import "github.com/gin-gonic/gin"

type IOrderController interface {
	NewOrder(*gin.Context)
}

type orderController struct {
}

func NewRouter() IOrderController {
	return &orderController{}
}

// ----------------------------------

func (*orderController) NewOrder(ctx *gin.Context) {

}
