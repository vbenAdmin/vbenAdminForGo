package main

import (
	"fmt"
	"goVben/Redis"
	"goVben/Routers"
	"goVben/Utils/Config"
	"goVben/Utils/Jwt"
)

func main() {
	jwt, err2 := Jwt.GenerateJWT(11, Config.Conf.JwtKey)
	if err2 != nil {
		println("出错了")
		return
	}
	println(jwt, "加密成功")
	claims, err3 := Jwt.ParseJwt(jwt, Config.Conf.JwtKey)
	if err3 != nil {
		fmt.Println("parse jwt failed, ", err3)
		return
	}
	fmt.Printf("%+v\n", claims.UserId)
	fmt.Printf("%+v\n", claims.ExpiresAt)
	fmt.Printf("%+v\n", claims.IssuedAt)
	err2 = Redis.Init()
	if err2 != nil {
		fmt.Println("redis启动失败", err2.Error())
	}
	r := Routers.SetupRouter()
	// Listen and Server n 0.0.0.0:8080
	err := r.Run(":39874")
	if err != nil {

	}

}
