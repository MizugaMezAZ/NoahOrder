package service

import (
	"gorder/api/repository"
	"gorder/model"
	"gorder/util/uuid"
	"time"
)

type IBillService interface {
	CreateBill(args model.Bill) error
}

type billService struct {
	or repository.IBillRepository
}

func NewBillService(or repository.IBillRepository) IBillService {
	return &billService{
		or: or,
	}
}

// ----------------------------------
func (o *billService) CreateBill(b model.Bill) error {
	id := uuid.GenUUID()

	now := time.Now()

	b.ID = id.ID
	b.EncodeID = id.Base58
	b.CreatedTime = now
	b.ExpirationTime = now.Add(model.ExpirationTime)

	return o.or.CreateBill(b)
}
