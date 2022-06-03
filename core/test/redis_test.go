package test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     "175.178.94.46:6379",
	Password: "lazy@ZYX1234!", // no password set
	DB:       0,               // use default DB
})

func TestRedis(t *testing.T) {

	err := rdb.Set("1", "2", time.Second*1000).Err()
	if err != nil {
		log.Fatal(err)
	}
}
