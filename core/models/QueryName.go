package models

import (
	"cloud-disk/core/helper"
	"errors"
)

func QueryName(name string, password string) (error, string) {
	user := new(UserBasic)
	// 1、从数据库中查询当前用户
	has, err := Engine.Where("name=? and password =?", name, helper.Md5(password)).Get(user)
	if err != nil {
		return errors.New(error.Error(err)), ""
	}
	if !has {
		return errors.New("用户名或密码错误"), ""
	}
	// 2、生成token
	s, err := helper.GetToken(uint64(user.Id), user.Identity, user.Name)
	if err != nil {
		return nil, err.Error()
	}
	return err, s
}
