package models

import (
	"cloud-disk/core/helper"
)

func QueryName(name string, password string) string {
	user := new(UserBasic)
	// 1、从数据库中查询当前用户
	has, err := Engine.Where("name=? and password =?", name, helper.Md5(password)).Get(user)
	if err != nil {
		return err.Error()
	}
	if !has {
		return err.Error()
	}
	// 2、生成token
	s, err := helper.GetToken(uint64(user.Id), user.Identity, user.Name)
	if err != nil {
		return err.Error()
	}
	return s
}
