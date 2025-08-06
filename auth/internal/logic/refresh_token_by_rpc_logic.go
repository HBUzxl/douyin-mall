package logic

import (
	"context"
	"fmt"

	"github.com/HBUzxl/douyin-mall/auth/auth"
	"github.com/HBUzxl/douyin-mall/auth/internal/svc"
	"github.com/HBUzxl/douyin-mall/common/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenByRpcLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshTokenByRpcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenByRpcLogic {
	return &RefreshTokenByRpcLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 刷新token
func (l *RefreshTokenByRpcLogic) RefreshTokenByRpc(in *auth.RefreshTokenReq) (*auth.DeliverTokenResp, error) {
	// todo: add your logic here and delete this line
	publicKeyString := l.svcCtx.Config.JWT.PublicSecret
	tokenExpire := l.svcCtx.Config.JWT.TokenExpire
	refreshTokenExpire := l.svcCtx.Config.JWT.RefreshTokenExpire
	privateKeyString := l.svcCtx.Config.JWT.PrivateSecret

	uuid, isRt, err := utils.VertifyToken(in.RefreshToken, publicKeyString)
	if err != nil {
		return nil, err
	}

	logx.Info("isRt: ", isRt)
	if !isRt {
		return nil, fmt.Errorf("Not a Refresh Token")
	}

	token, refreshToken, err := utils.SignToken(uuid, privateKeyString, tokenExpire, refreshTokenExpire)
	if err != nil {
		return nil, err
	}

	return &auth.DeliverTokenResp{
		Token:                   token,
		RefreshToken:            refreshToken,
		TokenExpireAfter:        tokenExpire,
		RefreshTokenExpireAfter: refreshTokenExpire,
	}, nil
}
