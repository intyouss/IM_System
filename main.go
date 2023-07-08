package main

import (
	"IM_System/router"
	"IM_System/utils"
)

func main() {
	utils.InitSystem()
	r := router.Router()
	err := r.Run(":8080")
	if err != nil {
		panic("项目启动发生错误")
	}
}
