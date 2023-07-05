package service

import (
	"IM_System/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var GetSwagger = func() gin.HandlerFunc {
	docs.SwaggerInfo.BasePath = ""
	return ginSwagger.WrapHandler(swaggerFiles.Handler)
}()
