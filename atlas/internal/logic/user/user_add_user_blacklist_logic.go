package user

import (
	"context"
	"errors"

	"github.com/HBUzxl/douyin-mall/atlas/internal/dal"
	"github.com/HBUzxl/douyin-mall/atlas/internal/svc"
	"github.com/HBUzxl/douyin-mall/atlas/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAddUserBlacklistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加用户黑名单
func NewUserAddUserBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddUserBlacklistLogic {
	return &UserAddUserBlacklistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAddUserBlacklistLogic) UserAddUserBlacklist(req *types.AddUserBlacklistReq) (resp *types.AddUserBlacklistResp, err error) {
	// todo: add your logic here and delete this line
	ok, err := l.svcCtx.CasbinEnforcer.AddRoleForUser(req.UserUuid, dal.BANNED_ROLE)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("add user blacklist failed")
	}

	return &types.AddUserBlacklistResp{
		Ok: true,
	}, nil
}
