package repository

import (
	"gorder/model"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type IOrderRepository interface {
	CreateOrder(model.Order) error
}

type orderRepository struct {
	db  *sqlx.DB
	rdb *redis.Client
}

func NewOrderRepository(db *sqlx.DB, rdb *redis.Client) IOrderRepository {
	return &orderRepository{
		db:  db,
		rdb: rdb,
	}
}

// ----------------------------------

func (o *orderRepository) CreateOrder(args model.Order) error {
	return nil
}
