package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UUID        string `gorm:"type:char(36);uniqueIndex"`
	UserUuid    string `gorm:"not null"`
	ProductUuid string `gorm:"not null"`
	Quantity    int64  `gorm:"not null"`
}
