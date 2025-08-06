package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Swagger Swagger
	UserRpc zrpc.RpcClientConf
	AuthRpc zrpc.RpcClientConf
}

type Swagger struct {
	Host     string `json:"host"`
	IsEnable bool   `json:"IsEnable"`
}
