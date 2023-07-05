package db

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func NewDB() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true&loc=UTC",
		"root",
		"1234",
		"127.0.0.1",
		"backend")

	DB, err = sqlx.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("connect server failed, err:%v\n", err)
	}

	DB.SetMaxOpenConns(200)
	DB.SetMaxIdleConns(30)
	DB.SetConnMaxLifetime(5 * time.Minute)
}
