package test

import (
	"bytes"
	"cloud-disk/core/define"
	"context"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func TestCloudcos(t *testing.T) {
	u, _ := url.Parse(define.Url)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: define.SECRETID,
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: define.SecretKey,
		},
	})

	key := "cloud-disk/lcetron.jpg"

	_, _, err := client.Object.Upload(
		context.Background(), key, "../img/lcetronlog.jpg", nil,
	)
	if err != nil {
		panic(err)
	}
}
func TestCloud(t *testing.T) {
	u, _ := url.Parse(define.Url)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: define.SECRETID,
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: define.SecretKey,
		},
	})

	key := "cloud-disk/lcetron111.jpg"
	b2, err := os.ReadFile("../img/lcetronlog.jpg")
	if err != nil {
		panic(err)
	}
	_, err = client.Object.Put(
		context.Background(), key, bytes.NewReader(b2), nil,
	)
	if err != nil {
		panic(err)
	}
}
