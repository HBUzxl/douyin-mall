package logic

import (
	"context"

	"github.com/HBUzxl/douyin-mall/common/utils"
	"github.com/HBUzxl/douyin-mall/product/internal/dao/model"
	"github.com/HBUzxl/douyin-mall/product/internal/svc"
	"github.com/HBUzxl/douyin-mall/product/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductLogic {
	return &CreateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateProductLogic) CreateProduct(in *product.CreateProductReq) (*product.CreateProductResp, error) {
	nodeId := l.svcCtx.Config.NodeID
	uuid, err := utils.GenerateSnowflakeID(nodeId)
	if err != nil {
		return nil, err
	}

	newProduct := &model.Product{
		Uuid:        uuid,
		Name:        in.Name,
		Description: in.Description,
		Price:       in.Price,
		Stock:       in.Stock,
	}
	result := l.svcCtx.DB.Create(newProduct)
	if result.Error != nil {
		return nil, result.Error
	}

	return &product.CreateProductResp{
		Product: &product.Product{
			Uuid:        newProduct.Uuid,
			Name:        newProduct.Name,
			Description: newProduct.Description,
			Price:       newProduct.Price,
			Stock:       newProduct.Stock,
		},
	}, nil
}
