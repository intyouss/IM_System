package service

import (
	"IM_System/models/db"
	"github.com/gin-gonic/gin"
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

// GetUser
// @Summary 获取单一用户
// @Tags 用户模块
// @param id query string true "id"
// @Success 200 {string} json "{"code","data", "msg"}"
// @Router /user/getUser [get]
func GetUser(c *gin.Context) {
	user := db.UserBasic{}
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"code": -1,
			"data": nil,
			"msg":  "服务端出现错误",
		})
		return
	}
	user.ID = uint(id)
	data, err := db.GetUser(&user)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
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
// @Success 200 {string} json "{"code","data", "msg"}"
// @Router /user/createUser [post]
func CreateUser(c *gin.Context) {
	user := db.UserBasic{}
	user.Name = c.PostForm("name")
	password := c.PostForm("password")
	repassword := c.PostForm("repassword")
	if password != repassword {
		c.JSON(403, gin.H{
			"code": -1,
			"data": nil,
			"msg":  "两次输入密码不一致",
		})
		return
	}
	user.Password = password
	err := db.CreateUser(&user)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
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
			"msg":  "服务端出现错误",
		})
		return
	}
	user.ID = uint(id)
	err = db.DeleteUser(&user)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
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
// @param name formData string false "name"
// @param password formData string false "password"
// @param Phone formData string false "Phone"
// @param Email formData string false "Email"
// @Success 200 {string} json "{"code","data", "msg"}"
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := db.UserBasic{}
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"code": -1,
			"data": nil,
			"msg":  "服务端出现错误",
		})
		return
	}
	user.ID = uint(id)
	user.Password = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Name = c.PostForm("name")
	user.Email = c.PostForm("email")
	err = db.UpdateUser(&user)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
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
