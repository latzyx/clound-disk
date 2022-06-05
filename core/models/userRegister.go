package models

import (
	"cloud-disk/core/helper"
	"errors"
	"log"
	"time"
)

func UserRegister(u *UserBasic, code string) (error, *UserBasic) {
	// 判断Code 验证码是否一致
	c, err := RDB.Get(u.Email).Result()
	log.Println("redis is err", c)
	log.Println("redis is err", u)
	log.Println("redis is err", code)
	if err != nil {
		return err, nil
	}
	if c != code {
		err = errors.New("验证码错误")
		return err, nil
	}
	// 判断用户名是否存在
	i, err := Engine.Where("name=?", u.Name).Count(new(UserBasic))
	if err != nil {
		return err, nil
	}
	if i > 0 {
		err = errors.New("用户名已存在")
		return err, nil
	}
	// 数据入库
	user := &UserBasic{
		Identity:  helper.GetUUID(),
		Name:      u.Name,
		Email:     u.Email,
		Password:  helper.Md5(u.Password),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	i2, err := Engine.Insert(user)
	if err != nil {
		return err, nil
	}
	log.Println(i2)
	return err, &UserBasic{}
}
