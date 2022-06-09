package models

import (
	"cloud-disk/core/define"
	"log"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var Engine = InInt()
var RDB = InintRedis()

func InInt() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", define.DataSource)
	if err != nil {
		log.Printf("Xrom New InInt Eroor:%v", err)
		return nil
	}
	return engine
}
func InintRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     define.MyredisAddr,
		Password: "000415", // no password set
		DB:       0,
	})
	return rdb
}
