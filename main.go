package main

import (
	"IM_System/router"
)

func main() {
	r := router.Router()
	err := r.Run()
	if err != nil {
		panic("项目启动发生错误")
	}
}
