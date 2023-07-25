package controller

import (
	"gorder/api/service"
	"gorder/logger"
	"gorder/model"
	"gorder/util/response"

	"github.com/gin-gonic/gin"
)

type IBillController interface {
	CreateBill(*gin.Context)
	GetBill(c *gin.Context)
}

type billController struct {
	bs service.IBillService
}

func NewBillController(bs service.IBillService) IBillController {
	return &billController{
		bs: bs,
	}
}

// ----------------------------------

func (bc *billController) CreateBill(c *gin.Context) {
	resp := response.Gin{Ctx: c}

	args := struct {
		Price     uint16 `json:"price"`
		PartySize uint8  `json:"party_size"`
		Area      string `json:"area"`
	}{}

	if err := c.ShouldBindJSON(&args); err != nil {
		resp.ErrResponse(err)
		return
	}

	data := model.Bill{
		Price:     args.Price,
		PartySize: args.PartySize,
		Area:      args.Area,
	}

	if err := bc.bs.CreateBill(&data); err != nil {
		logger.Error("創建新單失敗 Err: %v \n", err)
		resp.ErrResponse(err)
		return
	}

	logger.Infof("創建新單成功 ID:%d, encodeID:%s", data.ID, data.EncodeID)
	resp.Response(200, "成功", gin.H{
		"id": data.EncodeID,
	})

}

func (bc *billController) GetBill(c *gin.Context) {
	resp := response.Gin{Ctx: c}

	encodeid := c.Param("id")

	data, err := bc.bs.GetBill(encodeid)
	if err != nil {
		resp.ErrResponse(err)
		return
	}

	resp.Response(200, "成功", gin.H{
		"data": data,
	})
}
