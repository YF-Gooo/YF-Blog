package main

import (
	"singo/conf"
	"singo/server"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := server.NewRouter()
	r.Run(":3000")
	// 自签名https
	// r.RunTLS(":3000", "./cert/server.crt", "./cert/server.key")
}
