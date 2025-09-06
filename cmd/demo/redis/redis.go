package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := connRdb()

	err := rdb.Set(ctx, "Sid:admin", "test", 5*time.Second).Err()
	if err != nil {
		panic(err)
	}

	SID, err := rdb.Get(ctx, "Sid:admin").Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		panic(err)
	}

	fmt.Println(SID)
}

func connRdb() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	return rdb
}
