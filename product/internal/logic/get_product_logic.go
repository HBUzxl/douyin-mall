package logic

import (
	"context"
	"errors"

	"github.com/HBUzxl/douyin-mall/product/internal/dao/model"
	"github.com/HBUzxl/douyin-mall/product/internal/svc"
	"github.com/HBUzxl/douyin-mall/product/product"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductLogic {
	return &GetProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取商品
func (l *GetProductLogic) GetProduct(in *product.GetProductReq) (*product.GetProductResp, error) {
	// todo: add your logic here and delete this line
	dbProduct := &model.Product{}
	result := l.svcCtx.DB.Where("uuid = ?", in.Uuid).First(dbProduct)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.New("product not found")
		}
		return nil, result.Error
	}
	return &product.GetProductResp{
		Product: &product.Product{
			Uuid:        dbProduct.Uuid,
			Name:        dbProduct.Name,
			Price:       dbProduct.Price,
			Stock:       dbProduct.Stock,
			Description: dbProduct.Description,
		},
	}, nil
}
