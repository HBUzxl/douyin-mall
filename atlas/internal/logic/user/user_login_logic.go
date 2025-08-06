package user

import (
	"context"
	"fmt"

	"github.com/HBUzxl/douyin-mall/atlas/internal/svc"
	"github.com/HBUzxl/douyin-mall/atlas/internal/types"
	"github.com/HBUzxl/douyin-mall/auth/auth"
	"github.com/HBUzxl/douyin-mall/common/errorx"
	"github.com/HBUzxl/douyin-mall/user/user"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登录
func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line

	// 调用用户服务验证登录
	loginResp, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, errorx.NewErrCode(errorx.ErrInternalServerError)
	}

	if loginResp.UserUuid == "" {
		return nil, errorx.NewErrCode(errorx.INVALID_ACCOUNT_PASSWORD_ERROR)
	}

	uuid := loginResp.UserUuid

	// 调用auth服务获取token
	deliveryTokenResp, err := l.svcCtx.AuthRpc.DeliverTokenByRpc(l.ctx, &auth.DeliverTokenReq{
		UserUuid: uuid,
	})
	if err != nil {
		return nil, errorx.NewErrCode(errorx.AUTH_DELIBER_TOKEN_ERROR)
	}

	// 将token存入Redis
	redisKey := fmt.Sprintf("token: %s", deliveryTokenResp.Token)
	err = l.svcCtx.Redis.Setex(redisKey, "", int(deliveryTokenResp.TokenExpireAfter))
	if err != nil {
		return nil, errorx.NewErrCode(errorx.UNKNOWN_SERVER_ERROR)
	}

	return &types.LoginResp{
		Token:        deliveryTokenResp.Token,
		RefreshToken: deliveryTokenResp.RefreshToken,
	}, nil
}
