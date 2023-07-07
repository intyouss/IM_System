package models

import (
	"IM_System/utils"
	"gorm.io/gorm"
)

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

func SearchFriends(userID uint) []UserBasic {
	contacts := make([]UserContact, 0)
	objIDs := make([]uint, 0)
	utils.DB.Where("owner_id = ? and type = 1", userID).Find(&contacts)
	for _, v := range contacts {
		objIDs = append(objIDs, v.TargetId)
	}
	users := make([]UserBasic, 0)
	utils.DB.Where("id in ?", objIDs).Find(&users)
	return users
}
