package logic

import (
	"context"

	"github.com/HBUzxl/douyin-mall/product/internal/svc"
	"github.com/HBUzxl/douyin-mall/product/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductListLogic {
	return &GetProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取商品列表
func (l *GetProductListLogic) GetProductList(in *product.GetProductListReq) (*product.GetProductListResp, error) {
	// todo: add your logic here and delete this line

	db := l.svcCtx.DB.Table("product")

	if in.Name != nil {
		db = db.Where("name LIKE ?", "%"+*in.Name+"%")
	}
	if in.MinPrice != nil {
		db = db.Where("price >= ?", *in.MinPrice)
	}
	if in.MaxPrice != nil {
		db = db.Where("price <= ?", *in.MaxPrice)
	}

	var total int64
	products := make([]*product.Product, 0)
	res := db.Order("created_at DESC").
		Count(&total).
		Offset(int((in.Page - 1) * in.PageSize)).
		Limit(int(in.PageSize)).
		Find(&products)

	if res.Error != nil {
		return nil, res.Error
	}

	return &product.GetProductListResp{
		Products: products,
		Total:    total,
	}, nil
}
