// Package router 路由/*
package router

import (
	"IM_System/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", service.GetSwagger)
	r.GET("/index", service.GetIndex)
	r.GET("/user/getUserList", service.GetUserList)
	r.GET("/user/getUserOnly", service.GetUserOnly)
	r.POST("/user/createUser", service.CreateUser)
	r.DELETE("/user/deleteUser/:id", service.DeleteUser)
	r.PUT("/user/updateUser/:id", service.UpdateUser)
	r.POST("/user/userLogin", service.UserLogin)

	r.GET("/user/sendMsg", service.SendMsg)
	r.GET("/user/sendUserMsg", service.SendUserMsg)

	return r
}
