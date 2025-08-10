package user

import (
	"context"
	"errors"

	"github.com/HBUzxl/douyin-mall/atlas/internal/svc"
	"github.com/HBUzxl/douyin-mall/atlas/internal/types"
	"github.com/HBUzxl/douyin-mall/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserGetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewUserGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserGetUserInfoLogic {
	return &UserGetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserGetUserInfoLogic) UserGetUserInfo(req *types.GetUserInfoReq) (resp *types.GetUserInfoResp, err error) {
	// todo: add your logic here and delete this line
	// 从 context 获取 uuid
	uuid, ok := l.ctx.Value("uuid").(string)
	if !ok {
		return nil, errors.New("uuid not found in context")
	}

	// 调用远程 user 服务
	userResp, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &user.GetUserInfoReq{
		UserUuid: uuid,
	})
	if err != nil {
		return nil, err
	}

	// 使用 Casbin 获取用户角色
	roles, err := l.svcCtx.CasbinEnforcer.GetRolesForUser(uuid)
	if err != nil {
		return nil, err
	}

	// 返回响应
	return &types.GetUserInfoResp{
		Email: userResp.Email,
		Roles: roles,
	}, nil
}
