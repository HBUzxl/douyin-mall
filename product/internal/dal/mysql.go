package dal

import (
	"fmt"
	"time"

	"github.com/HBUzxl/douyin-mall/product/internal/config"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/test/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(c *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(c.MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("connect db %s failed, %w", c.MySQL.DSN, err)
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic(fmt.Errorf("auto migrate failed, %w", err))
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get raw sql db failed, %w", err)
	}

	sqlDB.SetMaxOpenConns(c.MySQL.MaxOpenConns)
	sqlDB.SetMaxIdleConns(c.MySQL.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(c.MySQL.ConnMaxLifetime) * time.Second)
	fmt.Println("mysql init success")
	return db, nil
}
