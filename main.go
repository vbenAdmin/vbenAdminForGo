package main

import (
	"fmt"
	"goVben/Redis"
	"goVben/Routers"
)

func main() {
	err2 := Redis.Init()
	if err2 != nil {
		fmt.Println("redis启动失败", err2.Error())
	}
	r := Routers.SetupRouter()
	// Listen and Server n 0.0.0.0:8080
	err := r.Run(":39874")
	if err != nil {

	}

}
