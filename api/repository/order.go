package repository

import (
	"context"
	"gorder/model"
	"time"

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

var ctx = context.Background()

// ----------------------------------

func (o *orderRepository) CreateOrder(args model.Order) error {
	query := `INSERT INTO 
					order(
						id,
						people,
						price,
						area,
						created_time,
						expiration_time
					)VALUES (
						:id,
						:people,
						:price,
						:area,
						:created_time,
						:expiration_time
					)`

	if _, err := o.db.NamedExec(query, args); err != nil {
		return err
	}

	o.rdb.HSet(ctx, "order", args.ID, args, 90*time.Minute)
	return nil
}

func (o *orderRepository) GetOrder(id string) (model.Order, error) {
	o.rdb.HGet(ctx, "order", id).Bytes()
	return model.Order{}, nil
}
