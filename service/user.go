package service

import (
	"IM_System/models"
	"IM_System/utils"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// GetUserList
// @Summary 获取用户列表
// @Tags 用户模块
// @Success 200 {string} json "{"code","data", "msg"}"
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
		"data": models.GetUserList(),
		"msg":  "success",
	})
}

// UserLogin
// @Summary 用户登录
// @Tags 用户模块
// @param name formData string true "用户名"
// @param password formData string true "密码"
// @Success 200 {string} json "{"code","data", "msg"}"
// @Router /user/userLogin [post]
func UserLogin(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	user, err := models.UserLogin(name, password)
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
		"data": user,
		"msg":  "success",
	})
}

// GetUserOnly
// @Summary 获取单一用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param phone query string false "手机号码"
// @param email query string false "邮箱"
// @Success 200 {string} json "{"code","data", "msg"}"
// @Router /user/getUserOnly [get]
func GetUserOnly(c *gin.Context) {
	user := models.UserBasic{}
	user.Phone = c.Query("phone")
	user.Name = c.Query("name")
	user.Email = c.Query("email")
	data, err := models.GetUserOnly(&user)
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
	user := models.UserBasic{}
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
	user.Salt = salt
	err := models.CreateUser(&user)
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
// @param id path string true "id"
// @Success 200 {string} json "{"code","data", "msg"}"
// @Router /user/deleteUser/{id} [delete]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"code": -1,
			"data": nil,
			"msg":  "Server error",
		})
		return
	}
	user.ID = uint(id)
	err = models.DeleteUser(&user)
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
// @param id path string true "id"
// @param name formData string false "用户名"
// @param password formData string false "密码"
// @param phone formData string false "手机号码"
// @param email formData string false "邮箱"
// @Success 200 {string} json "{"code","data", "msg"}"
// @Router /user/updateUser/{id} [put]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, err := strconv.Atoi(c.Param("id"))
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
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"data": nil,
			"msg":  "Parameter validation failed",
		})
		return
	}
	err = models.UpdateUser(&user)
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

// SearchFriends
// @Summary 搜索好友
// @Tags 用户模块
// @param userId formData string true "userId"
// @Success 200 {string} json "{"code","data", "msg"}"
// @Router /searchFriends [post]
func SearchFriends(c *gin.Context) {
	data, err := strconv.Atoi(c.PostForm("userId"))
	if err != nil {
		fmt.Println(err)
	}
	users := models.SearchFriends(uint(data))
	c.JSON(200, gin.H{
		"code": 0,
		"data": users,
		"msg":  "success",
	})
}

var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMsg(c *gin.Context) {
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)
	MsgHandler(ws, c)
}

func MsgHandler(ws *websocket.Conn, c *gin.Context) {
	go func() {
		ch := utils.Subscribe(c, utils.PublishKey)
		for msg := range ch {
			tm := time.Now().Format("2006-01-02 15:04:05")
			m := fmt.Sprintf("[ws][%s]:%s", tm, msg.Payload)
			err := ws.WriteMessage(1, []byte(m))
			if err != nil {
				log.Fatalln(err)
			}
		}
	}()
	for {
		_, data, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
		}
		utils.Publish(c, utils.PublishKey, string(data))
	}
}

func SendUserMsg(c *gin.Context) {
	models.Chat(c.Writer, c.Request)
}
