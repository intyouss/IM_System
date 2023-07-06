package db

import (
	"IM_System/utils"
	"errors"
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	Name          string
	Password      string
	Phone         string `valid:"matches(^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\\d{8}$)"`
	Email         string `valid:"email"`
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
	return data
}

func GetUser(user *UserBasic) (*UserBasic, error) {
	if err := utils.DB.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(user *UserBasic) (err error) {
	_, err = GetUser(&UserBasic{Name: user.Name})
	if err == nil {
		return errors.New("object is exist")
	}
	utils.DB.Create(&user)
	return nil
}

func DeleteUser(user *UserBasic) (err error) {
	_, err = GetUser(user)
	if err != nil {
		return err
	}
	utils.DB.Delete(&user)
	return nil
}

func UpdateUser(user *UserBasic) (err error) {
	data := UserBasic{}
	data.ID = user.ID
	_, err = GetUser(&data)
	if err != nil {
		return err
	}
	utils.DB.Model(&user).Updates(&UserBasic{
		Name: user.Name, Phone: user.Phone, Email: user.Email, Password: user.Password})
	return nil
}
