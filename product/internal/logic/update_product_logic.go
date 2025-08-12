package logic

import (
	"context"
	"errors"
	"fmt"

	"github.com/HBUzxl/douyin-mall/product/internal/dao/model"
	"github.com/HBUzxl/douyin-mall/product/internal/svc"
	"github.com/HBUzxl/douyin-mall/product/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductLogic {
	return &UpdateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新商品
func (l *UpdateProductLogic) UpdateProduct(in *product.UpdateProductReq) (*product.UpdateProductResp, error) {
	// todo: add your logic here and delete this line
	oldProduct := &model.Product{}
	result := l.svcCtx.DB.Where("uuid = ?", in.Product.Uuid).First(oldProduct)
	if result.Error != nil {
		return nil, result.Error
	}
	oldStock := oldProduct.Stock

	dbProduct := &model.Product{}
	updateResult := l.svcCtx.DB.Model(dbProduct).Where("uuid = ?", in.Product.Uuid).Updates(in.Product)
	if updateResult.Error != nil {
		return nil, updateResult.Error
	}
	if updateResult.RowsAffected == 0 {
		return nil, errors.New("product not found")
	}

	// 计算库存diff
	stockDiff := in.Product.Stock - oldStock
	if stockDiff != 0 {
		redisKey := fmt.Sprintf("product:stock:%s", in.Product.Uuid)

		// 检查Redis中是否存在该商品
		exists, err := l.svcCtx.Redis.Exists(redisKey)
		if err != nil {
			return nil, fmt.Errorf("failed to check redis key existence: %v", err)
		}
		if exists {
			var redisOp error
			_, redisOp = l.svcCtx.Redis.Incrby(redisKey, int64(stockDiff))
			if redisOp != nil {
				return nil, fmt.Errorf("failed to decrement redis key: %v", redisOp)
			}
		}
	}

	return &product.UpdateProductResp{
		Product: &product.Product{
			Uuid:        dbProduct.Uuid,
			Name:        dbProduct.Name,
			Price:       dbProduct.Price,
			Stock:       dbProduct.Stock,
			Description: dbProduct.Description,
		},
	}, nil
}
