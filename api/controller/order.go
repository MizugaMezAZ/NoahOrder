package controller

import "github.com/gin-gonic/gin"

type IOrder interface {
	NewOrder(*gin.Context)
}

type order struct {
}

func NewRouter() IOrder {
	return &order{}
}

// ----------------------------------

func (*order) NewOrder(ctx *gin.Context) {

}
