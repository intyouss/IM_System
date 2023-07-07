package models

import "gorm.io/gorm"

type UserGroup struct {
	gorm.Model
	Name    string
	OwnerId uint
	Icon    string
	Type    int
	Desc    string
}

func (table *UserGroup) TableName() string {
	return "user_group"
}
