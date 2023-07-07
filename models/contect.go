package models

import "gorm.io/gorm"

// UserContact 人员关系
type UserContact struct {
	gorm.Model
	OwnerId  uint //谁的关系信息
	TargetId uint //对应的谁 /群 ID
	Type     int  //对应的类型  1好友  2群  3xx
	Desc     string
}

func (table *UserContact) TableName() string {
	return "user_contact"
}
