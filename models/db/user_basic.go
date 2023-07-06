package db

import (
	"IM_System/utils"
	"errors"
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

func FindUserByName(name string) (user *UserBasic, err error) {
	db := utils.DB.Where("name = ?", name).First(&user)
	return user, db.Error
}

func FindUserByPhone(phone string) (user *UserBasic, err error) {
	db := utils.DB.Where("phone = ?", phone).First(&user)
	return user, db.Error
}

func FindUserByEmail(email string) (user *UserBasic, err error) {
	db := utils.DB.Where("email = ?", email).First(&user)
	return user, db.Error
}

func FindUserByID(id uint) (user *UserBasic, err error) {
	db := utils.DB.Where("id = ?", id).First(&user)
	return user, db.Error
}

func UserLogin(name string, password string) (user *UserBasic, err error) {
	user, err = FindUserByName(name)
	if err != nil {
		return nil, errors.New("username is not exist")
	}
	if passwd := utils.MakePassword(password, user.Salt); passwd != user.Password {
		return nil, errors.New("wrong username or password")
	}
	return user, nil
}

func GetUserOnly(user *UserBasic) (*UserBasic, error) {
	if user.Name != "" {
		return FindUserByName(user.Name)
	}
	if user.Phone != "" {
		return FindUserByPhone(user.Phone)
	}
	if user.Email != "" {
		return FindUserByEmail(user.Email)
	}
	return nil, errors.New("parameter error")
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	return data
}

func CreateUser(user *UserBasic) (err error) {
	if user.Name != "" {
		_, err = FindUserByName(user.Name)
		if err == nil {
			return errors.New("name is exist")
		}
	}
	if user.Email != "" {
		_, err = FindUserByEmail(user.Email)
		if err == nil {
			return errors.New("email is exist")
		}
	}
	if user.Phone != "" {
		_, err = FindUserByPhone(user.Phone)
		if err == nil {
			return errors.New("phone is exist")
		}
	}
	utils.DB.Create(&user)
	return nil
}

func DeleteUser(user *UserBasic) (err error) {
	_, err = FindUserByID(user.ID)
	if err != nil {
		return err
	}
	utils.DB.Delete(&user)
	return nil
}

func UpdateUser(user *UserBasic) (err error) {
	_, err = FindUserByID(user.ID)
	if err != nil {
		return err
	}
	utils.DB.Model(&user).Updates(&UserBasic{
		Name: user.Name, Phone: user.Phone, Email: user.Email, Password: user.Password})
	return nil
}
