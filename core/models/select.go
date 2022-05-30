package models

import "errors"

func Select(identity string) (*UserBasic, error) {
	ub := new(UserBasic)
	has, err := Engine.Where("identity=?", identity).Get(ub)
	if err != nil {
		return &UserBasic{}, err
	}
	if !has {
		return &UserBasic{}, errors.New("user is found")
	}
	return &UserBasic{}, err
}
