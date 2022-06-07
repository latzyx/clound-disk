package define

import "github.com/golang-jwt/jwt/v4"

type UserClaim struct {
	Id       uint64
	Identity string
	Name     string
	jwt.StandardClaims
}

var JwtKey = "cloud-disk-key"

var MailPassword = "jwficgnyahwsiefg"

var CodeExpire = 300

var SECRETID = "AKIDYAYidKp8D4g4D84nJbeS2NUMzNWu53mG"

var SecretKey = "IL9i391LYYLwKXr0veEeuztb9ZqwqX2N"

var Url = "https://1-1302719385.cos.ap-nanjing.myqcloud.com"
