package service

import (
	"gorder/api/repository"
	"gorder/model"
)

type IOrderService interface {
	CreateOrder(args model.Order) error
}

type orderService struct {
	or repository.IOrderRepository
}

func NewOrderService(or repository.IOrderRepository) IOrderService {
	return &orderService{
		or: or,
	}
}

// ----------------------------------
func (o *orderService) CreateOrder(args model.Order) error {
	return nil
}
