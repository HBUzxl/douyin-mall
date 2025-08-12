package product

import (
	"context"

	"github.com/HBUzxl/douyin-mall/atlas/internal/svc"
	"github.com/HBUzxl/douyin-mall/atlas/internal/types"
	"github.com/HBUzxl/douyin-mall/product/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取商品列表
func NewGetProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductListLogic {
	return &GetProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductListLogic) GetProductList(req *types.GetProductListReq) (resp *types.GetProductListResp, err error) {
	// todo: add your logic here and delete this line
	getProductListResp, err := l.svcCtx.ProductRpc.GetProductList(l.ctx, &product.GetProductListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Name:     &req.Name,
		MinPrice: &req.MinPrice,
		MaxPrice: &req.MaxPrice,
	})
	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)

	for _, product := range getProductListResp.Products {
		products = append(products, types.Product{
			Uuid:        product.Uuid,
			Name:        product.Name,
			Price:       product.Price,
			Stock:       product.Stock,
			Description: product.Description,
		})
	}

	return &types.GetProductListResp{
		Products: products,
		Total:    getProductListResp.Total,
	}, nil
}
