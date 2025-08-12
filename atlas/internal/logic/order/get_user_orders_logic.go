package order

import (
	"context"
	"errors"

	"github.com/HBUzxl/douyin-mall/atlas/internal/svc"
	"github.com/HBUzxl/douyin-mall/atlas/internal/types"
	"github.com/HBUzxl/douyin-mall/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserOrdersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户订单
func NewGetUserOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserOrdersLogic {
	return &GetUserOrdersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserOrdersLogic) GetUserOrders(req *types.GetUserOrdersReq) (resp *types.GetUserOrdersResp, err error) {
	// todo: add your logic here and delete this line
	userUuid, ok := l.ctx.Value("user_uuid").(string)
	if !ok {
		return nil, errors.New("user_uuid not found")
	}

	getUserOrdersResp, err := l.svcCtx.OrderRpc.GetUserOrders(l.ctx, &order.GetUserOrdersReq{
		UserUuid: userUuid,
	})
	if err != nil {
		return nil, err
	}

	orders := make([]types.Order, 0, len(getUserOrdersResp.Orders))
	for _, order := range getUserOrdersResp.Orders {
		items := make([]types.OrderItemWithPrice, 0, len(order.OrderItems))
		for _, item := range order.OrderItems {
			items = append(items, types.OrderItemWithPrice{
				ProductUuid: item.ProductUuid,
				Price:       item.Price,
				Quantity:    item.Quantity,
			})
		}

		orders = append(orders, types.Order{
			Uuid:        order.Uuid,
			UserUuid:    order.UserUuid,
			AddressUuid: order.AddressUuid,
			TotalPrice:  order.TotalPrice,
			Status:      order.Status,
			CreatedAt:   order.CreatedAt,
			OrderItems:  items,
		})
	}

	return &types.GetUserOrdersResp{
		Orders: orders,
		Total:  getUserOrdersResp.Total,
	}, nil
}
