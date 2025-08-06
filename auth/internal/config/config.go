package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	JWT struct {
		PrivateSecret      string
		PublicSecret       string
		TokenExpire        int64
		RefreshTokenExpire int64
	}

	Redis struct {
		Host string
		Type string
		Pass string
	}
}
