// Package router 路由/*
package router

import (
	"IM_System/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	//swagger
	r.GET("/swagger/*any", service.GetSwagger)

	//静态资源
	r.Static("/asset", "asset/")
	r.StaticFile("/favicon.ico", "asset/images/favicon.ico")
	r.LoadHTMLGlob("views/**/*")

	//首页
	r.GET("/index", service.GetIndex)
	r.GET("/toRegister", service.ToRegister)
	r.GET("/toChat", service.ToChat)
	r.GET("/chat", service.Chat)
	r.POST("/searchFriends", service.SearchFriends)

	//用户模块
	r.GET("/user/getUserList", service.GetUserList)
	r.GET("/user/getUserOnly", service.GetUserOnly)
	r.POST("/user/createUser", service.CreateUser)
	r.DELETE("/user/deleteUser/:id", service.DeleteUser)
	r.PUT("/user/updateUser/:id", service.UpdateUser)
	r.POST("/user/userLogin", service.UserLogin)

	//发送消息
	r.GET("/user/sendMsg", service.SendMsg)
	r.GET("/user/sendUserMsg", service.SendUserMsg)

	return r
}
