package model

import "time"

type Order struct {
	ID             uint64    `json:"id"`
	People         uint8     `json:"people"`
	Price          uint16    `json:"price"`
	Area           string    `json:"area"`
	CreatedTime    time.Time `json:"created_time" db:"created_time"`
	ExpirationTime time.Time `json:"expiration_time" db:"expiration_time"`
}

type UUID struct {
	ID     int64
	Base58 string
}
