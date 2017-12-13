package main

import "github.com/ajian/cloudgo-start/server"

func main() {
	// 端口为8080
	port := ":8080"

	// 取一个新的服务器实例并监听8080端口
	server := server.NewServer()
	server.Run(port)
}
