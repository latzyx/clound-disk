package logic

import (
	"context"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepostitorySaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepostitorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepostitorySaveLogic {
	return &UserRepostitorySaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepostitorySaveLogic) UserRepostitorySave(req *types.UserRepostitorySaveRequest) (resp *types.UserRepostitorySaveReply, err error) {
	// todo: add your logic here and delete this line

	return
}
