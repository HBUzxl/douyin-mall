package logic

import (
	"context"

	"github.com/HBUzxl/douyin-mall/common/utils"
	"github.com/HBUzxl/douyin-mall/user/internal/dao/model"
	"github.com/HBUzxl/douyin-mall/user/internal/svc"
	"github.com/HBUzxl/douyin-mall/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户注册
func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	// todo: add your logic here and delete this line
	nodeId := l.svcCtx.Config.NodeID
	uuid, err := utils.GenerateSnowflakeID(nodeId)
	if err != nil {
		return nil, err
	}

	newUser := &model.User{
		UUID:     uuid,
		Email:    in.Email,
		Password: utils.MD5Crypto(in.Password, l.svcCtx.Config.MD5Secret),
	}
	if err := l.svcCtx.DB.Create(newUser).Error; err != nil {
		return nil, err
	}
	return &user.RegisterResp{
		UserUuid: uuid,
	}, nil
}
