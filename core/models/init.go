package models

import (
	_ `github.com/go-sql-driver/mysql`
	`log`
	`xorm.io/xorm`
)

var Engine = InInt()

func InInt() *xorm.Engine {
	s := "root:107827@Zyx@tcp(127.0.0." +
		"1:3306)/cloud-disk?charset=utf8mb4&parseTime=True&loc=Local"
	engine, err := xorm.NewEngine("mysql", s)
	if err != nil {
		log.Printf("Xrom New InInt Eroor:%v", err)
		return nil
	}
	return engine
}
