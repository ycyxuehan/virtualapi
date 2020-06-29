package main

import (
	"fmt"

	"bing89.com/virtualapi/admin"
	"bing89.com/virtualapi/libs"
	"github.com/gin-gonic/gin"
)

func main(){
	fmt.Println("加载配置...")
	err := libs.Config.Load("config.json")
	if err != nil {
		//panic(err)
	}
	//启动管理后台
	router := gin.New()
	group := router.Group("/admin")
	group.GET("/service/:name", admin.GetService)
	group.POST("/api", admin.AddAPI)
	group.POST("/service", admin.AddService)
	group.POST("/group", admin.AddGroup)
	group.GET("/api/:name/methods", admin.GetMethods)
	group.GET("/services", admin.GetServices)
	go router.Run("0.0.0.0:2999")

	done := make(chan libs.Message)
	running := make(map[string]bool)
	for _, s := range libs.Config.Services {
		go s.Run(done)
	}
	for {
		select {
			case msg := <- done: {
				if msg.Running {
					running[msg.Service] = msg.Running
				}else {
					fmt.Println(msg.Service, "已停止运行。")
					delete(running, msg.Service)
				}
				if len(running) == 0 {
					libs.Config.Save("config.json")
				}
			}
		}
	}
}