package logic

import (
	"context"
	"errors"

	"github.com/HBUzxl/douyin-mall/product/internal/dao/model"
	"github.com/HBUzxl/douyin-mall/product/internal/svc"
	"github.com/HBUzxl/douyin-mall/product/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductLogic {
	return &DeleteProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除商品
func (l *DeleteProductLogic) DeleteProduct(in *product.DeleteProductReq) (*product.DeleteProductResp, error) {
	// todo: add your logic here and delete this line
	result := l.svcCtx.DB.Where("uuid = ?", in.Uuid).Delete(&model.Product{})
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("product not found")
	}
	return &product.DeleteProductResp{
		Uuid: in.Uuid,
	}, nil
}
