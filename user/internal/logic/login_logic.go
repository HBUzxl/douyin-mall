package logic

import (
	"context"
	"errors"

	"github.com/HBUzxl/douyin-mall/common/utils"
	"github.com/HBUzxl/douyin-mall/user/internal/dao/model"
	"github.com/HBUzxl/douyin-mall/user/internal/svc"
	"github.com/HBUzxl/douyin-mall/user/user"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	// todo: add your logic here and delete this line
	cryptoPassword := utils.MD5Crypto(in.Password, l.svcCtx.Config.MD5Secret)
	loginUser := &model.User{
		Email:    in.Email,
		Password: cryptoPassword,
	}

	if err := l.svcCtx.DB.Where(loginUser).First(loginUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &user.LoginResp{
				UserUuid: "",
			}, nil
		}
		return nil, err
	}
	return &user.LoginResp{
		UserUuid: loginUser.UUID,
	}, nil
}
