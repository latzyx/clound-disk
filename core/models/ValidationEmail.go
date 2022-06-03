package models

import (
	"cloud-disk/core/define"
	"cloud-disk/core/helper"
	"errors"
	"time"
)

func ValiDationEmail(email string) (string, error) {
	i, err := Engine.Where("email=?", email).Count(new(UserBasic))
	if err != nil {
		return "", err
	}
	if i > 0 {
		err = errors.New("该邮箱已被注册")
		return "", err
	}
	code := helper.MailCode()
	RDB.Set(email, code, time.Second*time.Duration(define.CodeExpire))

	return code, err
}
