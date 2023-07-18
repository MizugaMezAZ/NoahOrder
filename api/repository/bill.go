package repository

import (
	"context"
	"errors"
	"gorder/logger"
	"gorder/model"
	"gorder/util/json"
	"gorder/util/response"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type IBillRepository interface {
	CreateBill(model.Bill) error
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

func (b *billRepository) CreateBill(data model.Bill) error {
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

	if err := b.rdb.HSet(ctx, model.RedisBillHash, data.ID, string(bt)).Err(); err != nil {
		err = errors.New("redis存帳單資料失敗 Err: " + err.Error())
		return err
	}

	return nil
}

func (b *billRepository) GetBill(id string) (data model.Bill, err error) {
	bt, err := b.rdb.HGet(ctx, model.RedisBillHash, id).Bytes()
	if err != nil {
		err = response.Custom("帳單不存在")
		return
	}

	if err = json.Unmarshal(bt, &data); err != nil {
		err = errors.New("redis解析帳單資料失敗 Err: " + err.Error())
		return
	}

	if data.ExpirationTime.After(time.Now()) {
		err = b.rdb.HDel(ctx, model.RedisBillHash, id).Err()
		if err != nil {
			logger.Warnf("bill過期刪除失敗 id:%s \n", id)
		}

		err = response.Custom("帳單不存在")
		return
	}

	return
}
