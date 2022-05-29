package logic

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.LonginRequest) (resp *types.LoginReply, err error) {
	// todo: add your logic here and delete this line
	user:= new(models.UserBasic)
	 // 1、从数据库中查询当前用户
	has, err := models.Engine.Where("name=? and password =?", req.Name, helper.Md5(req.Password)).Get(user)
	if err != nil {
		return nil ,err	
	}
	if !has {
			return nil,errors.New("用户名或密码错误")
	}
	// 2、生成token
	s, err := helper.GetToken(uint64(user.Id), user.Identity, user.Name)
	if err != nil {
		return nil, err
	}
	resp=new(types.LoginReply)
	resp.Token=s
	return
}
