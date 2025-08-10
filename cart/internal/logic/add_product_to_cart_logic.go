package logic

import (
	"context"
	"errors"

	"github.com/HBUzxl/douyin-mall/cart/cart"
	"github.com/HBUzxl/douyin-mall/cart/internal/dao/model"
	"github.com/HBUzxl/douyin-mall/cart/internal/svc"
	"github.com/HBUzxl/douyin-mall/common/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProductToCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductToCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductToCartLogic {
	return &AddProductToCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加商品到购物车
func (l *AddProductToCartLogic) AddProductToCart(in *cart.AddProductToCartReq) (*cart.AddProductToCartResp, error) {
	// todo: add your logic here and delete this line
	cartItem := &model.Cart{}
	findRes := l.svcCtx.DB.
		Where("user_uuid = ?", in.Item.UserUuid).
		Where("product_uuid = ?", in.Item.ProductUuid).
		First(cartItem)

	if findRes.Error != nil {
		return nil, findRes.Error
	}

	if findRes.RowsAffected == 0 {
		if in.Item.Quantity <= 0 {
			return nil, errors.New("quantity should be positive when createing")
		}
		nodeId := l.svcCtx.Config.NodeID
		uuid, err := utils.GenerateSnowflakeID(nodeId)
		if err != nil {
			return nil, err
		}

		createCart := &model.Cart{
			UUID:        uuid,
			UserUuid:    in.Item.UserUuid,
			ProductUuid: in.Item.ProductUuid,
			Quantity:    in.Item.Quantity,
		}
		createResp := l.svcCtx.DB.Create(createCart)
		if createResp.Error != nil {
			return nil, createResp.Error
		}

		return &cart.AddProductToCartResp{
			Item: &cart.CartItem{
				UserUuid:    in.Item.UserUuid,
				ProductUuid: in.Item.ProductUuid,
				Quantity:    in.Item.Quantity,
			},
		}, nil
	}

	afterQuantity := cartItem.Quantity + in.Item.Quantity
	if afterQuantity <= 0 {
		deleteResp := l.svcCtx.DB.Delete(cartItem)
		if deleteResp.Error != nil {
			return nil, deleteResp.Error
		}

		return &cart.AddProductToCartResp{
			Item: &cart.CartItem{
				UserUuid:    in.Item.UserUuid,
				ProductUuid: in.Item.ProductUuid,
				Quantity:    afterQuantity,
			},
		}, nil
	}

	cartItem.Quantity = afterQuantity
	updateResp := l.svcCtx.DB.Save(cartItem)
	if updateResp.Error != nil {
		return nil, updateResp.Error
	}
	if updateResp.RowsAffected == 0 {
		return nil, errors.New("!!! Product may be deleted by other service")
	}

	return &cart.AddProductToCartResp{
		Item: &cart.CartItem{
			UserUuid:    in.Item.UserUuid,
			ProductUuid: in.Item.ProductUuid,
			Quantity:    afterQuantity,
		},
	}, nil
}
