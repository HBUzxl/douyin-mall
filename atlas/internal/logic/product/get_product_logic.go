package product

import (
	"context"

	"github.com/HBUzxl/douyin-mall/atlas/internal/svc"
	"github.com/HBUzxl/douyin-mall/atlas/internal/types"
	"github.com/HBUzxl/douyin-mall/product/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取商品
func NewGetProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductLogic {
	return &GetProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductLogic) GetProduct(req *types.GetProductReq) (resp *types.GetProductResp, err error) {
	// todo: add your logic here and delete this line
	getResp, err := l.svcCtx.ProductRpc.GetProduct(l.ctx, &product.GetProductReq{
		Uuid: req.Uuid,
	})
	if err != nil {
		return nil, err
	}
	return &types.GetProductResp{
		Product: types.Product{
			Uuid:        getResp.Product.Uuid,
			Name:        getResp.Product.Name,
			Price:       getResp.Product.Price,
			Stock:       getResp.Product.Stock,
			Description: getResp.Product.Description,
		},
	}, nil
}
