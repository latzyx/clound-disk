package models

import (
	"log"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var Engine = InInt()
var RDB = InintRedis()

func InInt() *xorm.Engine {
	s := "lazylast:lazy@ZYX1234!@tcp(175.178.94.46:3306)/cloud_disk?charset=utf8mb4&parseTime=True&loc=Local"
	engine, err := xorm.NewEngine("mysql", s)
	if err != nil {
		log.Printf("Xrom New InInt Eroor:%v", err)
		return nil
	}
	return engine
}
func InintRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "175.178.94.46:6379",
		Password: "000415", // no password set
		DB:       0,
	})
	return rdb
}
