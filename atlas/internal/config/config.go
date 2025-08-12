package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Swagger    Swagger
	UserRpc    zrpc.RpcClientConf
	AuthRpc    zrpc.RpcClientConf
	CartRpc    zrpc.RpcClientConf
	ProductRpc zrpc.RpcClientConf
	OrderRpc   zrpc.RpcClientConf

	Redis struct {
		Host string
		Type string
		Pass string
	}

	MySQL struct {
		DSN             string
		MaxOpenConns    int
		MaxIdleConns    int
		ConnMaxLifetime int
	}

	CasbinModel string
}

type Swagger struct {
	Host     string `json:"host"`
	IsEnable bool   `json:"IsEnable"`
}
