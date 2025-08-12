package svc

import (
	"fmt"

	"github.com/HBUzxl/douyin-mall/product/internal/config"
	"github.com/HBUzxl/douyin-mall/product/internal/dal"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := dal.NewDB(&c)
	if err != nil {
		panic(fmt.Sprintf("new database failed, %v", err))
	}
	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
	}
}
