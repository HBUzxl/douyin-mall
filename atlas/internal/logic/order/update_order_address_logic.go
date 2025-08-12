package order

import (
	"context"
	"errors"

	"github.com/HBUzxl/douyin-mall/atlas/internal/svc"
	"github.com/HBUzxl/douyin-mall/atlas/internal/types"
	"github.com/HBUzxl/douyin-mall/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新订单地址
func NewUpdateOrderAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderAddressLogic {
	return &UpdateOrderAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateOrderAddressLogic) UpdateOrderAddress(req *types.UpdateOrderAddressReq) (resp *types.UpdateOrderAddressResp, err error) {
	// todo: add your logic here and delete this line
	userUuid, ok := l.ctx.Value("user_uuid").(string)
	if !ok {
		return nil, errors.New("user_uuid not found")
	}

	_, err = l.svcCtx.OrderRpc.UpdateOrderAddress(l.ctx, &order.UpdateOrderAddressReq{
		UserUuid:    userUuid,
		OrderUuid:   req.OrderUuid,
		AddressUuid: req.AddressUuid,
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdateOrderAddressResp{
		Success: true,
	}, nil
}
