package dal

import (
	"github.com/HBUzxl/douyin-mall/atlas/internal/config"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

const (
	ADMIN_ROLE    = "Admin"
	CUSTOMER_ROLE = "Customer"
	SELLER_ROLE   = "Seller"
	BANNED_ROLE   = "Banned"
)

const (
	CUSTOMER_OBJECT = "customer_obj"
	SELLER_OBJECT   = "seller_obj"
)

var Enforcer *casbin.Enforcer

func InitCasbin(c config.Config, db gorm.DB) (*casbin.Enforcer, error) {
	adapter, err := gormadapter.NewAdapterByDB(&db)
	if err != nil {
		return nil, err
	}

	// 创建Enforcer
	enforcer, err := casbin.NewEnforcer(c.CasbinModel, adapter)
	if err != nil {
		return nil, err
	}

	enforcer.EnableAutoSave(true)

	// 加载策略
	err = enforcer.LoadPolicy()
	if err != nil {
		return nil, err
	}

	// 初始化默认策略
	addInitPolicy(enforcer)

	return enforcer, nil
}

func addInitPolicy(enforcer *casbin.Enforcer) {
	// 添加角色权限
	if _, err := enforcer.HasPolicy(ADMIN_ROLE, ".*"); err != nil {
		enforcer.AddPolicy(ADMIN_ROLE, ".*")
	}
	if _, err := enforcer.HasPolicy(CUSTOMER_ROLE, CUSTOMER_OBJECT); err != nil {
		enforcer.AddPolicy(CUSTOMER_ROLE, CUSTOMER_OBJECT)
	}
	if _, err := enforcer.HasPolicy(SELLER_ROLE, SELLER_OBJECT); err != nil {
		enforcer.AddPolicy(SELLER_ROLE, SELLER_OBJECT)
	}
	if _, err := enforcer.HasPolicy(SELLER_ROLE, CUSTOMER_OBJECT); err != nil {
		enforcer.AddPolicy(SELLER_ROLE, CUSTOMER_OBJECT)
	}

	// 绑定超级管理员用户（UUID = "0000000000000000000"）
	if _, err := enforcer.HasRoleForUser("0000000000000000000", ADMIN_ROLE); err != nil {
		enforcer.AddRoleForUser("0000000000000000000", ADMIN_ROLE)
	}

	// 保存到 DB
	_ = enforcer.SavePolicy()
}
