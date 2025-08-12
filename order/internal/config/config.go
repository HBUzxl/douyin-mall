package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	MySQL struct {
		DSN             string
		MaxOpenConns    int
		MaxIdleConns    int
		ConnMaxLifetime int
	}
}
