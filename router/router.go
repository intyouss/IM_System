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
	return r
}
