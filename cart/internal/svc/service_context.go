package svc

import (
	"fmt"

	"github.com/HBUzxl/douyin-mall/cart/internal/config"
	"github.com/HBUzxl/douyin-mall/cart/internal/dal"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := dal.NewDB(&c)
	if err != nil {
		panic(fmt.Sprintf("new database failed, %v", err))
	}
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
