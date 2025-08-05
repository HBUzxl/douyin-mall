package svc

import (
	"douyin-mall/atlas/internal/config"
	"douyin-mall/user/userclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {

	var userRpc userclient.User
	userRpc = userclient.NewUser(zrpc.MustNewClient(c.UserRpc))

	return &ServiceContext{
		Config:  c,
		UserRpc: userRpc,
	}
}
