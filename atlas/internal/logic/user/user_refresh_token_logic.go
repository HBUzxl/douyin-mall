package user

import (
	"context"
	"fmt"

	"github.com/HBUzxl/douyin-mall/atlas/internal/svc"
	"github.com/HBUzxl/douyin-mall/atlas/internal/types"
	"github.com/HBUzxl/douyin-mall/auth/auth_client"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 刷新token
func NewUserRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRefreshTokenLogic {
	return &UserRefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRefreshTokenLogic) UserRefreshToken(req *types.RefreshTokenReq) (resp *types.RefreshTokenResp, err error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.AuthRpc.RefreshTokenByRpc(l.ctx, &auth_client.RefreshTokenReq{
		RefreshToken: req.RefreshToken,
	})
	if err != nil {
		return nil, err
	}

	redisKey := fmt.Sprintf("token:%s", rpcResp.Token)
	err = l.svcCtx.Redis.Set(redisKey, rpcResp.Token)
	if err != nil {
		return nil, err
	}

	return &types.RefreshTokenResp{
		Token:        rpcResp.Token,
		RefreshToken: rpcResp.RefreshToken,
	}, nil
}
