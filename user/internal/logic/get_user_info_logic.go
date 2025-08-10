package logic

import (
	"context"

	"github.com/HBUzxl/douyin-mall/user/internal/dao/model"
	"github.com/HBUzxl/douyin-mall/user/internal/svc"
	"github.com/HBUzxl/douyin-mall/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户信息
func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	// todo: add your logic here and delete this line
	userInfo := &model.User{}
	result := l.svcCtx.DB.Where("uuid = ?", in.UserUuid).First(userInfo)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user.GetUserInfoResp{
		Email: userInfo.Email,
	}, nil
}
