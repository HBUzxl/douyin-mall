package cart

import (
	"context"
	"errors"

	"github.com/HBUzxl/douyin-mall/atlas/internal/svc"
	"github.com/HBUzxl/douyin-mall/atlas/internal/types"
	"github.com/HBUzxl/douyin-mall/cart/cart"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClearCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 清空购物车
func NewClearCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClearCartLogic {
	return &ClearCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClearCartLogic) ClearCart(req *types.ClearCartReq) (resp *types.ClearCartResp, err error) {
	// todo: add your logic here and delete this line
	userUuid, ok := l.ctx.Value("uuid").(string)
	if !ok {
		return nil, errors.New("user uuid not found")
	}

	deletResp, err := l.svcCtx.CartRpc.ClearCart(l.ctx, &cart.ClearCartReq{
		UserUuid: userUuid,
	})
	if err != nil {
		return nil, err
	}

	return &types.ClearCartResp{
		UserUuid: deletResp.UserUuid,
	}, nil
}
