package cart

import (
	"context"
	"errors"

	"github.com/HBUzxl/douyin-mall/atlas/internal/svc"
	"github.com/HBUzxl/douyin-mall/atlas/internal/types"
	"github.com/HBUzxl/douyin-mall/cart/cart_client"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProductToCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加商品到购物车
func NewAddProductToCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductToCartLogic {
	return &AddProductToCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddProductToCartLogic) AddProductToCart(req *types.AddProductToCartReq) (resp *types.AddProductToCartResp, err error) {
	// todo: add your logic here and delete this line
	userUuid, ok := l.ctx.Value("uuid").(string)
	if !ok {
		return nil, errors.New("user uuid not found")
	}
	grpcResp, err := l.svcCtx.CartRpc.AddProductToCart(l.ctx, &cart_client.AddProductToCartReq{
		Item: &cart_client.CartItem{
			UserUuid:    userUuid,
			ProductUuid: req.Item.ProductUuid,
			Quantity:    req.Item.Quantity,
		},
	})

	if err != nil {
		return nil, err
	}

	return &types.AddProductToCartResp{
		Item: types.CartItem{
			ProductUuid: grpcResp.Item.ProductUuid,
			Quantity:    grpcResp.Item.Quantity,
		},
	}, nil
}
