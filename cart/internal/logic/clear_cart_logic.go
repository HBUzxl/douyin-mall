package logic

import (
	"context"

	"github.com/HBUzxl/douyin-mall/cart/cart"
	"github.com/HBUzxl/douyin-mall/cart/internal/dao/model"
	"github.com/HBUzxl/douyin-mall/cart/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClearCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClearCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClearCartLogic {
	return &ClearCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 清空购物车
func (l *ClearCartLogic) ClearCart(in *cart.ClearCartReq) (*cart.ClearCartResp, error) {
	// todo: add your logic here and delete this line
	resp := l.svcCtx.DB.Where("user_uuid = ?", in.UserUuid).Delete(&model.Cart{})
	if resp.Error != nil {
		return nil, resp.Error
	}

	return &cart.ClearCartResp{
		UserUuid: in.UserUuid,
	}, nil
}
