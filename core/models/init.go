package models

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var Engine = InInt()

func InInt() *xorm.Engine {
	s := "root:107827@zyx@tcp(127.0.0." +
		"1:3306)/cloud_disk?charset=utf8mb4&parseTime=True&loc=Local"
	engine, err := xorm.NewEngine("mysql", s)
	if err != nil {
		log.Printf("Xrom New InInt Eroor:%v", err)
		return nil
	}
	return engine
}
