package logic

import (
	"context"
	"fmt"

	"github.com/HBUzxl/douyin-mall/auth/auth"
	"github.com/HBUzxl/douyin-mall/auth/internal/svc"
	"github.com/HBUzxl/douyin-mall/common/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeliverTokenByRpcLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeliverTokenByRpcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeliverTokenByRpcLogic {
	return &DeliverTokenByRpcLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分发token
func (l *DeliverTokenByRpcLogic) DeliverTokenByRpc(in *auth.DeliverTokenReq) (*auth.DeliverTokenResp, error) {
	// todo: add your logic here and delete this line
	tokenExpire := l.svcCtx.Config.JWT.TokenExpire
	refreshTokenExpire := l.svcCtx.Config.JWT.RefreshTokenExpire
	privateKeyString := l.svcCtx.Config.JWT.PrivateSecret

	// 解码私钥
	token, refreshToken, err := utils.SignToken(in.UserUuid, privateKeyString, tokenExpire, refreshTokenExpire)
	if err != nil {
		return nil, fmt.Errorf("SignToken failed: %v", err)
	}

	return &auth.DeliverTokenResp{
		Token:                   token,
		RefreshToken:            refreshToken,
		TokenExpireAfter:        tokenExpire,
		RefreshTokenExpireAfter: refreshTokenExpire,
	}, nil
}
