package svc

import (
	"github.com/HBUzxl/douyin-mall/atlas/internal/config"
	"github.com/HBUzxl/douyin-mall/atlas/internal/dal"
	"github.com/HBUzxl/douyin-mall/auth/auth_client"
	"github.com/HBUzxl/douyin-mall/cart/cart_client"
	"github.com/HBUzxl/douyin-mall/order/order_client"
	"github.com/HBUzxl/douyin-mall/product/product_client"
	"github.com/HBUzxl/douyin-mall/user/user_client"
	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config

	UserRpc    user_client.User
	AuthRpc    auth_client.Auth
	CartRpc    cart_client.Cart
	ProductRpc product_client.ProductZrpcClient
	OrderRpc   order_client.Order

	Db    *gorm.DB
	Redis *redis.Redis

	CasbinEnforcer *casbin.Enforcer
}

func NewServiceContext(c config.Config) *ServiceContext {

	userRpc := user_client.NewUser(zrpc.MustNewClient(c.UserRpc))
	authRpc := auth_client.NewAuth(zrpc.MustNewClient(c.AuthRpc))
	cartRpc := cart_client.NewCart(zrpc.MustNewClient(c.CartRpc))
	productRpc := product_client.NewProductZrpcClient(zrpc.MustNewClient(c.ProductRpc))
	orderRpc := order_client.NewOrder(zrpc.MustNewClient(c.OrderRpc))

	db, err := dal.NewDB(&c)
	if err != nil {
		panic(err)
	}

	casbinEnforcer, err := dal.InitCasbin(c, *db)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,

		UserRpc:    userRpc,
		AuthRpc:    authRpc,
		CartRpc:    cartRpc,
		ProductRpc: productRpc,
		OrderRpc:   orderRpc,

		Redis: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),

		CasbinEnforcer: casbinEnforcer,
	}
}
