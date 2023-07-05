package service

import (
	"IM_System/models"
	"IM_System/models/db"
	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	data := models.Message{
		Data: db.GetUserList(),
		Msg:  "success",
	}

	c.JSON(200, data)
}
