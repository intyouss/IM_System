package service

import (
	"IM_System/models/db"
	"IM_System/utils"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
)

// GetUserList
// @Summary 获取用户列表
// @Tags 用户模块
// @Success 200 {string} json "{"code","data", "msg"}"
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
		"data": db.GetUserList(),
		"msg":  "success",
	})
}

// GetUserOnly
// @Summary 获取单一用户
// @Tags 用户模块
// @param id query string false "id"
// @param name query string false "用户名"
// @param phone query string false "手机号码"
// @param email query string false "邮箱"
// @Success 200 {string} json "{"code","data", "msg"}"
// @Router /user/getUserOnly [get]
func GetUserOnly(c *gin.Context) {
	user := db.UserBasic{}
	if c.Query("id") != "" {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(500, gin.H{
				"code": -1,
				"data": nil,
				"msg":  "Server error",
			})
			return
		}
		user.ID = uint(id)
	}
	user.Phone = c.Query("phone")
	user.Name = c.Query("name")
	user.Email = c.Query("email")
	data, err := db.GetUserOnly(&user)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"data": nil,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"data": data,
		"msg":  "success",
	})
}

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @param name formData string true "用户名"
// @param password formData string true "密码"
// @param repassword formData string true "确认密码"
// @param phone formData string false "手机号码"
// @param email formData string false "邮箱"
// @Success 200 {string} json "{"code","data", "msg"}"
// @Router /user/createUser [post]
func CreateUser(c *gin.Context) {
	user := db.UserBasic{}
	user.Name = c.PostForm("name")
	password := c.PostForm("password")
	repassword := c.PostForm("repassword")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	salt := fmt.Sprintf("%06d", rand.Int31())
	if password != repassword {
		c.JSON(200, gin.H{
			"code": 1,
			"data": nil,
			"msg":  "Entered passwords differ",
		})
		return
	}
	user.Password = utils.MakePassword(password, salt)
	err := db.CreateUser(&user)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"data": nil,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"data": nil,
		"msg":  "success",
	})
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @param id formData string true "id"
// @Success 200 {string} json "{"code","data", "msg"}"
// @Router /user/deleteUser [post]
func DeleteUser(c *gin.Context) {
	user := db.UserBasic{}
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"code": -1,
			"data": nil,
			"msg":  "Server error",
		})
		return
	}
	user.ID = uint(id)
	err = db.DeleteUser(&user)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"data": nil,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"data": nil,
		"msg":  "success",
	})
}

// UpdateUser
// @Summary 更新用户
// @Tags 用户模块
// @param id formData string true "id"
// @param name formData string false "用户名"
// @param password formData string false "密码"
// @param phone formData string false "手机号码"
// @param email formData string false "邮箱"
// @Success 200 {string} json "{"code","data", "msg"}"
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := db.UserBasic{}
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"code": -1,
			"data": nil,
			"msg":  "Server error",
		})
		return
	}
	user.ID = uint(id)
	user.Password = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Name = c.PostForm("name")
	user.Email = c.PostForm("email")
	_, err = govalidator.ValidateStruct(user)
	fmt.Println(err)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"data": nil,
			"msg":  "Parameter validation failed",
		})
		return
	}
	err = db.UpdateUser(&user)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"data": nil,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"data": nil,
		"msg":  "success",
	})
}
