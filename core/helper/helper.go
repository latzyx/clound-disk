package helper

import (
	"context"
	"crypto/md5"
	"crypto/tls"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"net/smtp"
	"net/url"
	"path"
	"time"

	"cloud-disk/core/define"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/jordan-wright/email"
	"github.com/tencentyun/cos-go-sdk-v5"
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

// AnalyzeToken token 解码
func AnalyzeToken(token string) (*define.UserClaim, error) {
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(
		token, uc, func(token *jwt.Token) (interface{}, error) {
			return []byte(define.JwtKey), nil
		},
	)
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("token is invalid")
	}
	return uc, err
}

// 邮箱验证码发送

func MailSendCode(mail string, code string) error {
	e := email.NewEmail()
	e.From = "Get <name-zyx@foxmail.com>"
	e.To = []string{"2669738224@qq.com"}
	e.Subject = "验证码发送"
	e.HTML = []byte("你的验证码为:<h1>" + code + "</h1>")
	err := e.SendWithTLS(
		"smtp.qq.com:465",
		smtp.PlainAuth("", "name-zyx@foxmail.com", define.MailPassword, "smtp.qq.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"},
	)
	if err != nil {
		return err
	}
	return nil
}

func MailCode() string {
	s := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(100000))
	return s
}
func GetUUID() string {
	return uuid.New().String()
}

func CosUpload(r *http.Request) (string, error) {
	u, _ := url.Parse(define.Url)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(
		b, &http.Client{
			Transport: &cos.AuthorizationTransport{
				// 通过环境变量获取密钥
				// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
				SecretID: define.SECRETID,
				// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
				SecretKey: define.SecretKey,
			},
		},
	)
	file, fileHeader, err := r.FormFile("file")
	key := "cloud-disk/" + GetUUID() + path.Ext(fileHeader.Filename)
	_, err = client.Object.Put(context.Background(), key, file, nil)
	if err != nil {
		panic(err)
	}
	return define.Url + "/" + key, nil

}
