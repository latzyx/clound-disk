package test

import (
	"crypto/tls"
	"net/smtp"
	"testing"

	"github.com/jordan-wright/email"
)

func TestSendMail(t *testing.T) {
	e := email.NewEmail()
	e.From = "Get <name-zyx@foxmail.com>"
	e.To = []string{"2669738224@qq.com"}
	e.Subject = "验证码发送"
	e.HTML = []byte("你的验证码为:<h1>123456</h1>")
	err := e.SendWithTLS("smtp.qq.com:465", smtp.PlainAuth("", "name-zyx@foxmail.com", "jwficgnyahwsiefg", "smtp.qq.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
	if err != nil {
		t.Fatal(err)
	}
}
