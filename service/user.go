package service

import (
	"IM_System/models"
	"IM_System/models/db"
	"github.com/gin-gonic/gin"
)

// GetUserList
// @Tags 获取用户列表
// @Success 200 {object} models.Message
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := models.Message{
		Data: db.GetUserList(),
		Msg:  "success",
	}

	c.JSON(200, data)
}
