package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UUID     string `gorm:"type:char(36);uniqueIndex"`
	Email    string `gorm:"type:varchar(255);uniqueIndex"`
	Password string `gorm:"type:varchar(255)"`
}
