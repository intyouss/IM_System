package db

import (
	"IM_System/utils"
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	Name          string
	Password      string
	Phone         string
	Email         string
	ClientIp      string
	ClientPort    string
	Identity      string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LogoutTime    time.Time
	IsLogout      bool
	DeviceInfo    string
	gorm.Model
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	//for _, v := range data {
	//	fmt.Println(v)
	//}
	return data
}

func CreateUser(user *UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}

func DeleteUser(user *UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}

func UpdateUser(user *UserBasic) *gorm.DB {
	return utils.DB.Updates(&user)
}
