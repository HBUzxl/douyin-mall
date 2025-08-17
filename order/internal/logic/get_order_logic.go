package logic

import (
	"context"

	"github.com/HBUzxl/douyin-mall/order/internal/dao/model"
	"github.com/HBUzxl/douyin-mall/order/internal/svc"
	"github.com/HBUzxl/douyin-mall/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取订单信息
func (l *GetOrderLogic) GetOrder(in *order.GetOrderReq) (*order.GetOrderResp, error) {
	// todo: add your logic here and delete this line
	// 获取订单信息
	var createOrder model.Order
	orderResp := l.svcCtx.DB.Model(&model.Order{}).Where("uuid = ?", in.OrderUuid).First(&createOrder)
	if orderResp.Error != nil {
		return nil, orderResp.Error
	}

	// 获取订单商品信息
	var orderItems []*model.OrderItem
	orderItemsResp := l.svcCtx.DB.Model(&model.OrderItem{}).Where("order_uuid = ?", in.OrderUuid).Find(&orderItems)
	if orderItemsResp.Error != nil {
		return nil, orderItemsResp.Error
	}

	// 组装订单商品信息
	getOrderItemsResp := make([]*order.OrderItem, 0)
	for _, orderItem := range orderItems {
		getOrderItemsResp = append(getOrderItemsResp, &order.OrderItem{
			ProductUuid: orderItem.ProductUuid,
			Price:       orderItem.Price,
			Quantity:    orderItem.Quantity,
		})
	}

	return &order.GetOrderResp{
		Order: &order.OrderInfo{
			Uuid:       createOrder.UUID,
			UserUuid:   createOrder.UserUuid,
			TotalPrice: createOrder.Total,
			Status:     int64(createOrder.Status),
			CreatedAt:  createOrder.CreatedAt.Unix(),
			OrderItems: getOrderItemsResp,
		},
	}, nil
}
