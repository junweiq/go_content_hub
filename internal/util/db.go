package util

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

func ConnDB() *sqlx.DB {
	return sqlx.MustConnect("mysql", "root:handsome@tcp(localhost:3306)/go_cms?charset=utf8mb4&parseTime=True&loc=Local")
}

func ConnRdb() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	return rdb
}
