package logic

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(req *types.FileUploadRequest) (resp *types.FileUploadReply, err error) {
	// todo: add your logic here and delete this line
	rp := &models.RepositoryPool{
		Identity: helper.GetUUID(),
		Hash:     req.Hash,
		Name:     req.Name,
		Ext:      req.Ext,
		Path:     req.Path,
	}
	fmt.Println(rp.Identity)
	_, err = models.Engine.Insert(rp)
	if err != nil {
		return nil, err
	}
	resp = new(types.FileUploadReply)
	fmt.Println(resp.Identity)
	resp.Identity = rp.Identity
	fmt.Println("resp", resp.Identity)
	fmt.Println(resp)
	return
}
