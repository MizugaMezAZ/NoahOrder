package repository

import (
	"context"
	"errors"
	"gorder/model"
	"gorder/util/json"
	"gorder/util/lib"
	"gorder/util/response"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type IBillRepository interface {
	InsertNewBill(model.Bill) error
	GetBill(id string) (data model.Bill, err error)
}

type billRepository struct {
	db  *sqlx.DB
	rdb *redis.Client
}

func NewBillRepository(db *sqlx.DB, rdb *redis.Client) IBillRepository {
	return &billRepository{
		db:  db,
		rdb: rdb,
	}
}

var ctx = context.Background()

// ----------------------------------

func (b *billRepository) InsertNewBill(data model.Bill) error {
	query := `INSERT INTO 
					order(
						id,
						encode_id,
						price,
						party_size,
						area,
						created_time,
						expiration_time
					)VALUES (
						:id,
						:encode_id
						:price,
						:party_size,
						:area,
						:created_time,
						:expiration_time
					)`

	if _, err := b.db.NamedExec(query, data); err != nil {
		return err
	}

	bt, _ := json.Marshal(data)

	if err := b.rdb.Set(ctx, lib.RedisFormatKey(model.RedisBillHash, data.EncodeID), string(bt), model.ExpirationTime).Err(); err != nil {
		err = errors.New("redis存帳單資料失敗 Err: " + err.Error())
		return err
	}

	return nil
}

func (b *billRepository) GetBill(id string) (data model.Bill, err error) {
	bt, err := b.rdb.Get(ctx, lib.RedisFormatKey(model.RedisBillHash, id)).Bytes()
	if err != nil {
		err = response.Custom("帳單不存在")
		return
	}

	if err = json.Unmarshal(bt, &data); err != nil {
		err = errors.New("redis解析帳單資料失敗 Err: " + err.Error())
		return
	}

	return
}
