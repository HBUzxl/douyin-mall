package cart

import (
	"context"
	"errors"

	"github.com/HBUzxl/douyin-mall/atlas/internal/svc"
	"github.com/HBUzxl/douyin-mall/atlas/internal/types"
	"github.com/HBUzxl/douyin-mall/cart/cart"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取购物车
func NewGetCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCartLogic {
	return &GetCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCartLogic) GetCart(req *types.GetCartReq) (resp *types.GetCartResp, err error) {
	// todo: add your logic here and delete this line
	userUuid, ok := l.ctx.Value("uuid").(string)
	if !ok {
		return nil, errors.New("user uuid not found")
	}

	cartItems, err := l.svcCtx.CartRpc.GetCart(l.ctx, &cart.GetCartReq{
		UserUuid: userUuid,
	})
	if err != nil {
		return nil, err
	}

	items := make([]types.CartItem, 0)
	for _, item := range cartItems.CartItems {
		items = append(items, types.CartItem{
			ProductUuid: item.ProductUuid,
			Quantity:    item.Quantity,
		})
	}

	return &types.GetCartResp{
		Total:     int64(len(items)),
		CartItems: items,
	}, nil
}
