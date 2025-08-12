package model

import "gorm.io/gorm"

const (
	OrderStatusUnpaid    = 0
	OrderStatusPaid      = 1
	OrderStatusCancelled = -1
)

type Order struct {
	gorm.Model
	UUID        string `gorm:"type:char(36);uniqueIndex"`
	UserUuid    string `gorm:"type:char(36);"`
	AddressUuid string `gorm:"type:char(36);"`
	Total       int64  `gorm:"type:bigint;not null"`
	Status      int    `gorm:"type:int;not null"`
}

type OrderItem struct {
	gorm.Model
	UUID        string `gorm:"type:char(36);uniqueIndex"`
	OrderUUID   string `gorm:"type:char(36);"`
	ProductUuid string `gorm:"type:char(36);"`
	Price       int64  `gorm:"type:bigint;not null"`
	Quantity    int64  `gorm:"type:bigint;not null"`
}
