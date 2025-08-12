package product

import (
	"context"

	"github.com/HBUzxl/douyin-mall/atlas/internal/svc"
	"github.com/HBUzxl/douyin-mall/atlas/internal/types"
	"github.com/HBUzxl/douyin-mall/product/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新商品
func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductLogic {
	return &UpdateProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProductLogic) UpdateProduct(req *types.UpdateProductReq) (resp *types.UpdateProductResp, err error) {
	// todo: add your logic here and delete this line
	updateResp, err := l.svcCtx.ProductRpc.UpdateProduct(l.ctx, &product.UpdateProductReq{
		Product: req.Product,
	})
	if err != nil {
		return nil, err
	}
	return &types.UpdateProductResp{
		Product: types.Product{
			Uuid:        updateResp.Product.Uuid,
			Name:        updateResp.Product.Name,
			Price:       updateResp.Product.Price,
			Stock:       updateResp.Product.Stock,
			Description: updateResp.Product.Description,
		},
	}, nil
}
