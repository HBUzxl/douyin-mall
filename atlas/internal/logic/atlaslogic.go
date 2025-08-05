package logic

import (
	"context"

	"atlas/internal/svc"
	"atlas/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AtlasLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAtlasLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AtlasLogic {
	return &AtlasLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AtlasLogic) Atlas(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
