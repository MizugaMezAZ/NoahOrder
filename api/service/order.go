package service

import "gorder/api/repository"

type IOrderService interface {
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
