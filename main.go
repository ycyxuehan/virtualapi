package main

import (
	"fmt"
	"bing89.com/virtualapi/libs"
)

func main(){
	fmt.Println("加载配置...")
	config := libs.Configuration{}
	err := config.Load("config.json")
	if err != nil {
		panic(err)
	}
	done := make(chan libs.Message)
	running := make(map[string]bool)
	for _, s := range config.Services {
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
			}
		}
	}
}