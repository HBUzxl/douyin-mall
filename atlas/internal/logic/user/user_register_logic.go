package user

import (
	"context"

	"github.com/HBUzxl/douyin-mall/atlas/internal/dal"
	"github.com/HBUzxl/douyin-mall/atlas/internal/svc"
	"github.com/HBUzxl/douyin-mall/atlas/internal/types"
	"github.com/HBUzxl/douyin-mall/common/errorx"
	"github.com/HBUzxl/douyin-mall/user/user_client"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户注册
func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line

	if req.Password != req.ConfirmPassword {
		return nil, errorx.NewErrCode(errorx.PASSWORD_NOT_MATCH_ERROR)
	}

	rpcResp, err := l.svcCtx.UserRpc.Register(l.ctx, &user_client.RegisterReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, errorx.NewErrCode(errorx.USER_REGISTER_ERROR)
	}

	// 添加初始化customer角色权限
	l.svcCtx.CasbinEnforcer.AddRoleForUser(rpcResp.UserUuid, dal.CUSTOMER_ROLE)

	return &types.RegisterResp{
		UserUuid: rpcResp.UserUuid,
	}, nil
}
