package svc

import (
	"douyin-mall/atlas/internal/config"
	"douyin-mall/user/user_client"

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
