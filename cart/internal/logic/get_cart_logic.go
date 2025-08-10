package logic

import (
	"context"

	"github.com/HBUzxl/douyin-mall/cart/cart"
	"github.com/HBUzxl/douyin-mall/cart/internal/dao/model"
	"github.com/HBUzxl/douyin-mall/cart/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCartLogic {
	return &GetCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取购物车
func (l *GetCartLogic) GetCart(in *cart.GetCartReq) (*cart.GetCartResp, error) {
	// todo: add your logic here and delete this line
	var getCarts []*model.Cart
	resp := l.svcCtx.DB.Where("user_uuid = ?", in.UserUuid).Find(&getCarts)
	if resp.Error != nil {
		return nil, resp.Error
	}

	var cartItems []*cart.CartItem
	for _, item := range getCarts {
		cartItems = append(cartItems, &cart.CartItem{
			UserUuid:    item.UserUuid,
			ProductUuid: item.ProductUuid,
			Quantity:    item.Quantity,
		})
	}

	return &cart.GetCartResp{
		CartItems: cartItems,
	}, nil
}
