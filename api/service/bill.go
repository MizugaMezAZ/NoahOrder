package service

import (
	"gorder/api/repository"
	"gorder/model"
	"gorder/util/uuid"
	"time"
)

type IBillService interface {
	CreateBill(args *model.Bill) error
	GetBill(encodeid string) (model.Bill, error)
}

type billService struct {
	br repository.IBillRepository
}

func NewBillService(br repository.IBillRepository) IBillService {
	return &billService{
		br: br,
	}
}

// ----------------------------------
func (o *billService) CreateBill(b *model.Bill) error {
	snowid := uuid.GenUUID()

	now := time.Now()

	b.ID = snowid.ID
	b.EncodeID = snowid.Base58
	b.CreatedTime = now
	b.ExpirationTime = now.Add(model.ExpirationTime)

	return o.br.InsertNewBill(*b)
}

func (b *billService) GetBill(encodeid string) (model.Bill, error) {
	return b.br.GetBill(encodeid)
}
