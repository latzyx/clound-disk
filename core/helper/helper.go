package helper

import (
	"cloud-disk/core/define"
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"math/rand"
	"net/smtp"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GetToken(id uint64, identity string, name string) (string, error) {
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	s, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return s, nil

}

// 邮箱验证码发送

func MailSendCode(mail, code string) error {
	e := email.NewEmail()
	e.From = "Get <name-zyx@foxmail.com>"
	e.To = []string{"2669738224@qq.com"}
	e.Subject = "验证码发送"
	e.HTML = []byte("你的验证码为:<h1>" + code + "</h1>")
	err := e.SendWithTLS("smtp.qq.com:465", smtp.PlainAuth("", "name-zyx@foxmail.com", define.MailPassword, "smtp.qq.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
	if err != nil {
		return err
	}
	return nil
}

func MailCode() string {
	s := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(100000))

	return s
}
