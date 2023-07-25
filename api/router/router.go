package router

import (
	"gorder/api/controller"
	"gorder/api/middleware"

	"github.com/gin-gonic/gin"
)

type RouterParam struct {
	AuthController controller.IAuthController
	BillController controller.IBillController
}

type HttpRouter struct {
	ac controller.IAuthController
	bc controller.IBillController
}

func NewRouter(r RouterParam) *HttpRouter {
	return &HttpRouter{
		ac: r.AuthController,
		bc: r.BillController,
	}
}

func (r *HttpRouter) SetupRoute(e *gin.Engine) {
	e.Use(middleware.CrosHandler())

	e.GET("/bill/:orderid", r.bc.GetBill)
	e.POST("/bill", r.bc.CreateBill)
}
