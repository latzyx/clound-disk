package models

import (
	"errors"
)

func UserRegister(email string, name string) (error, *UserBasic) {
	has, err := Engine.Where("name=? ", name).Count(new(UserBasic))
	if err != nil {
		return nil, &UserBasic{}
	}
	if has > 1 {
		return errors.New(" The user name has been registered"), &UserBasic{}
	}
	sql := "select id forma user_Basic order by id Desc"
	b, err := Engine.Query(sql)

	Engine.Insert("id=? name=? email=?", b, name, email)
	return err, &UserBasic{}
}
