package model

import "time"

type Bill struct {
	ID             int64     `json:"id"`
	EncodeID       string    `json:"encode_id"  db:"encode_id"`
	Price          uint16    `json:"price"`
	PartySize      uint8     `json:"party_size" db:"party_size"`
	Area           string    `json:"area"`
	CreatedTime    time.Time `json:"created_time"    db:"created_time"`
	ExpirationTime time.Time `json:"expiration_time" db:"expiration_time"` // 資料採用redis HASH存儲 HASH沒有單一過期系統 寫在裡面自己判斷
}

type SnowID struct {
	ID     int64
	Base58 string
}
