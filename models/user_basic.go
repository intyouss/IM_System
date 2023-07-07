package models

import (
	"IM_System/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	Name          string
	Password      string `json:"-"`
	Phone         string `valid:"matches(^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\\d{8}$)"`
	Email         string `valid:"email"`
	ClientIp      string
	ClientPort    string
	Identity      string
	Salt          string `json:"-"`
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

func FindUserByName(name string) (user *UserBasic, db *gorm.DB) {
	db = utils.DB.Where("name = ?", name).First(&user)
	return user, db
}

func FindUserByPhone(phone string) (user *UserBasic, db *gorm.DB) {
	db = utils.DB.Where("phone = ?", phone).First(&user)
	return user, db
}

func FindUserByEmail(email string) (user *UserBasic, db *gorm.DB) {
	db = utils.DB.Where("email = ?", email).First(&user)
	return user, db
}

func FindUserByID(id uint) (user *UserBasic, db *gorm.DB) {
	db = utils.DB.Where("id = ?", id).First(&user)
	return user, db
}

func UserLogin(name string, password string) (*UserBasic, error) {
	user, db := FindUserByName(name)
	if db.Error != nil {
		return nil, errors.New("username is not exist")
	}
	if passwd := utils.MakePassword(password, user.Salt); passwd != user.Password {
		return nil, errors.New("wrong username or password")
	}
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := utils.MD5Encode(str)
	db.Update("identity", temp)
	return user, nil
}

func GetUserOnly(user *UserBasic) (*UserBasic, error) {
	if user.Name != "" {
		data, db := FindUserByName(user.Name)
		return data, db.Error
	}
	if user.Phone != "" {
		data, db := FindUserByPhone(user.Phone)
		return data, db.Error
	}
	if user.Email != "" {
		data, db := FindUserByEmail(user.Email)
		return data, db.Error
	}
	return nil, errors.New("parameter error")
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	return data
}

func CreateUser(user *UserBasic) error {
	if user.Name == "" {
		return errors.New("username not entered")
	} else {
		_, db := FindUserByName(user.Name)
		if db.Error == nil {
			return errors.New("username is exist")
		}
	}
	if user.Email != "" {
		_, db := FindUserByEmail(user.Email)
		if db.Error == nil {
			return errors.New("email is exist")
		}
	}
	if user.Phone != "" {
		_, db := FindUserByPhone(user.Phone)
		if db.Error == nil {
			return errors.New("phone is exist")
		}
	}
	utils.DB.Create(&user)
	return nil
}

func DeleteUser(user *UserBasic) error {
	_, db := FindUserByID(user.ID)
	if db.Error != nil {
		return db.Error
	}
	utils.DB.Delete(&user)
	return nil
}

func UpdateUser(user *UserBasic) error {
	_, db := FindUserByID(user.ID)
	if db.Error != nil {
		return db.Error
	}
	utils.DB.Model(&user).Updates(&UserBasic{
		Name: user.Name, Phone: user.Phone, Email: user.Email, Password: user.Password})
	return nil
}
