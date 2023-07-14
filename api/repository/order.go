package repository

import (
	"context"
	"gorder/model"
	"gorder/util/lib"
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
					)VALUE (
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

	o.rdb.Set(ctx, lib.RedisFormatKey("order", args.ID), args, 90*time.Minute)
	return nil
}
