package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Swagger Swagger
}

type Swagger struct {
	Host     string `json:"host"`
	IsEnable bool   `json:"IsEnable"`
}
