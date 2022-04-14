package main

import (
	"fmt"
	"net"

	"atem-test/grpc_server/service"

	"google.golang.org/grpc"
)

func main() {
	// 初始化一个grpc对象
	serv := grpc.NewServer()
	pbTeacher := service.NewFindClassService()
	// 注册服务
	pbTeacher.RegisterFindClassServer(serv)
	pbHello := service.NewHelloService()
	pbHello.RegisterHelloServer(serv)
	// 设置监听
	listener, err := net.Listen("tcp", ":8999")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	//启动服务
	err = serv.Serve(listener)
	if err != nil {
		fmt.Println(err)
		return
	}
}
