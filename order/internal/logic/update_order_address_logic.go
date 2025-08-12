package logic

import (
	"context"
	"errors"

	"github.com/HBUzxl/douyin-mall/order/internal/dao/model"
	"github.com/HBUzxl/douyin-mall/order/internal/svc"
	"github.com/HBUzxl/douyin-mall/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderAddressLogic {
	return &UpdateOrderAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新订单地址
func (l *UpdateOrderAddressLogic) UpdateOrderAddress(in *order.UpdateOrderAddressReq) (*order.UpdateOrderAddressResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.DB.Model(&model.Order{}).
		Where("uuid = ?", in.OrderUuid).
		Where("user_uuid = ?", in.UserUuid).
		Update("address_uuid", in.AddressUuid)

	if err.Error != nil {
		return nil, errors.New("update order address failed")
	}

	if err.RowsAffected == 0 {
		return nil, errors.New("order not found")
	}

	return &order.UpdateOrderAddressResp{
		AddressUuid: in.AddressUuid,
	}, nil
}
