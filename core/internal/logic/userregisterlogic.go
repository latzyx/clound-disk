package logic

import (
	"context"
	"errors"
	"log"
	"time"

	"cloud-disk/core/helper"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterReply, err error) {
	// 判断Code 验证码是否一致
	code, err := models.RDB.Get(req.Mail).Result()
	if err != nil {
		return nil, err
	}
	if code != req.Code {
		err = errors.New("验证码错误")
		return
	}
	// 判断用户名是否存在
	i, err := models.Engine.Where("name=?", req.Name).Count(new(models.UserBasic))
	if err != nil {
		return nil, err
	}
	if i > 0 {
		err = errors.New("用户名已存在")
		return
	}
	// 数据入库
	user := &models.UserBasic{
		Identity:  helper.GetUUID(),
		Name:      req.Name,
		Email:     req.Mail,
		Password:  helper.Md5(req.Password),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	i2, err := models.Engine.Insert(user)
	if err != nil {
		return nil, err
	}
	log.Println(i2)
	return
}
