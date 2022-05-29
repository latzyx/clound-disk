package helper

import (
	"cloud-disk/core/define"
	"crypto/md5"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

func Md5(s string) string {
	return fmt.Sprintf("%x",md5.Sum([]byte(s)))
}


func GetToken(id uint64,identity string,name string) (string,error){
	uc := define.UserClaim{
		Id:  id,
		Identity: identity,
		Name: name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,uc)
	s, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return s,nil

}