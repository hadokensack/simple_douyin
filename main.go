package main

import (
	"fmt"
	"simple_douyin/config"
	"simple_douyin/router"
)

func main() {
	r := router.Init_Router()                          //启动路由
	err := r.Run(fmt.Sprintf(":%d", config.Info.Port)) // 默认端口号8080
	if err != nil {
		return
	}
}
