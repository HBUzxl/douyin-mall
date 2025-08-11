package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Uuid        string `gorm:"type:char(36);uniqueIndex"`
	Name        string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:text"`
	Price       int64  `gorm:"type:bigint"`
	Stock       int64  `gorm:"type:bigint"`
}
