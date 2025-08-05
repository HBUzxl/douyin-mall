package svc

import (
	"github.com/HBUzxl/douyin-mall/atlas/internal/config"
	"github.com/HBUzxl/douyin-mall/user/user_client"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user_client.User
}

func NewServiceContext(c config.Config) *ServiceContext {

	var userRpc user_client.User
	userRpc = user_client.NewUser(zrpc.MustNewClient(c.UserRpc))

	return &ServiceContext{
		Config:  c,
		UserRpc: userRpc,
	}
}
