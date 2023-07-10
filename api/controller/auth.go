package controller

import (
	"gorder/util/response"

	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

type authController struct {
}

func NewAuthController() IAuthController {
	return &authController{}
}

// ----------------------------------

func (a *authController) Login(ctx *gin.Context) {
	resp := response.Gin{Ctx: ctx}

	// args := struct {
	// 	Account string `binding:"require"`
	// 	Password string `binding:"require"`
	// }{}

	resp.Response(200, "ok", nil)
}

func (a *authController) Logout(ctx *gin.Context) {
	resp := response.Gin{Ctx: ctx}
	resp.Response(200, "ok", nil)
}
