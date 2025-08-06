package svc

import (
	"github.com/HBUzxl/douyin-mall/atlas/internal/config"
	"github.com/HBUzxl/douyin-mall/auth/auth_client"
	"github.com/HBUzxl/douyin-mall/user/user_client"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user_client.User
	AuthRpc auth_client.Auth

	Redis *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {

	userRpc := user_client.NewUser(zrpc.MustNewClient(c.UserRpc))
	authRpc := auth_client.NewAuth(zrpc.MustNewClient(c.AuthRpc))

	return &ServiceContext{
		Config:  c,
		UserRpc: userRpc,
		AuthRpc: authRpc,
		Redis: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
	}
}
