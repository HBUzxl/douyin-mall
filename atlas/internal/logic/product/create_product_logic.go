package product

import (
	"context"

	"github.com/HBUzxl/douyin-mall/atlas/internal/svc"
	"github.com/HBUzxl/douyin-mall/atlas/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建商品
func NewCreateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductLogic {
	return &CreateProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateProductLogic) CreateProduct(req *types.CreateProductReq) (resp *types.CreateProductResp, err error) {
	// todo: add your logic here and delete this line

	createResp, err := l.svcCtx.ProductRpc.CreateProduct(l.ctx, &product.CreateProductReq{
		Name:        req.Name,
		Price:       req.Price,
		Stock:       req.Stock,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.CreateProductResp{
		Product: types.Product{
			Uuid:        createResp.Product.Uuid,
			Name:        createResp.Product.Name,
			Price:       createResp.Product.Price,
			Stock:       createResp.Product.Stock,
			Description: createResp.Product.Description,
		},
	}
	return
}
