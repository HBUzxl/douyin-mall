package logic

import (
	"context"

	"github.com/HBUzxl/douyin-mall/order/internal/dao/model"
	"github.com/HBUzxl/douyin-mall/order/internal/svc"
	"github.com/HBUzxl/douyin-mall/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserOrdersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserOrdersLogic {
	return &GetUserOrdersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户订单
func (l *GetUserOrdersLogic) GetUserOrders(in *order.GetUserOrdersReq) (*order.GetUserOrdersResp, error) {
	// todo: add your logic here and delete this line

	orders := make([]*model.Order, 0)
	res := l.svcCtx.DB.Model(&model.Order{}).Where("user_uuid = ?", in.UserUuid).Find(&orders)
	if res.Error != nil {
		return nil, res.Error
	}

	if len(orders) == 0 {
		return &order.GetUserOrdersResp{
			Orders: make([]*order.OrderInfo, 0),
			Total:  0,
		}, nil
	}

	orderUuids := make([]string, 0, len(orders))
	for _, order := range orders {
		orderUuids = append(orderUuids, order.UUID)
	}

	orderItems := make([]*model.OrderItem, 0)
	res = l.svcCtx.DB.Model(&model.OrderItem{}).Where("order_uuid IN (?)", orderUuids).Find(&orderItems)
	if res.Error != nil {
		return nil, res.Error
	}

	orderItemsMap := make(map[string][]*model.OrderItem)
	for _, orderItem := range orderItems {
		orderItemsMap[orderItem.OrderUUID] = append(orderItemsMap[orderItem.OrderUUID], orderItem)
	}

	orderResp := make([]*order.OrderInfo, 0, len(orders))
	for _, ord := range orders {
		items := orderItemsMap[ord.UUID]
		respItems := make([]*order.OrderItem, 0, len(items))
		for _, item := range items {
			respItems = append(respItems, &order.OrderItem{
				ProductUuid: item.ProductUuid,
				Price:       item.Price,
				Quantity:    item.Quantity,
			})
		}

		orderResp = append(orderResp, &order.OrderInfo{
			Uuid:        ord.UUID,
			UserUuid:    ord.UserUuid,
			AddressUuid: ord.AddressUuid,
			TotalPrice:  ord.Total,
			Status:      int64(ord.Status),
			CreatedAt:   ord.CreatedAt.Unix(),
			OrderItems:  respItems,
		})
	}

	return &order.GetUserOrdersResp{
		Orders: orderResp,
		Total:  int64(len(orders)),
	}, nil
}
