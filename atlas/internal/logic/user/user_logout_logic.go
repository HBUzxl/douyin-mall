package user

import (
	"context"
	"errors"

	"github.com/HBUzxl/douyin-mall/atlas/internal/svc"
	"github.com/HBUzxl/douyin-mall/atlas/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户退出
func NewUserLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogoutLogic {
	return &UserLogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogoutLogic) UserLogout(req *types.LogoutReq) (resp *types.LogoutResp, err error) {
	// todo: add your logic here and delete this line
	token, ok := l.ctx.Value("token").(string)
	if !ok {
		return nil, errors.New("token is not found")
	}

	// 从Redis删除token
	redisKey := "token:" + token
	_, err = l.svcCtx.Redis.Del(redisKey)
	if err != nil {
		return nil, err
	}

	return &types.LogoutResp{
		Ok: true,
	}, nil
}
