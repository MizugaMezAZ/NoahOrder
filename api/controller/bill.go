package controller

import "github.com/gin-gonic/gin"

type IBillController interface {
	CreateBill(*gin.Context)
}

type billController struct {
}

func NewRouter() IBillController {
	return &billController{}
}

// ----------------------------------

func (*billController) CreateBill(ctx *gin.Context) {

}
