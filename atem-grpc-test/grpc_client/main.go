package main

import (
	"atem-test/pbs/hello"
	"atem-test/pbs/student"
	"bufio"
	"context"
	"log"
	"os"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:8999", grpc.WithInsecure())
	if err != nil {
		log.Fatal("连接 gPRC 服务失败,", err)
	}

	defer conn.Close()

	go testFindClass(conn)
	testHello(conn)
}

func testFindClass(conn *grpc.ClientConn) {
	// 创建 gRPC 客户端
	grpcClient := student.NewFindClassClient(conn)
	// 创建请求参数
	req := &student.Student{Name: "songzb", Age: 13}
	// 发送请求，调用 MyTest 接口
	response, err := grpcClient.MyClass(context.Background(), req)
	if err != nil {
		log.Fatal("发送请求失败，原因是:", err)
	}
	log.Println(response)
}

func testHello(conn *grpc.ClientConn) {
	// 创建 gRPC 客户端
	grpcClient := hello.NewMsgHelloServiceClient(conn)
	// 创建请求参数
	req := &hello.RequestMsg{Name: "songzb"}
	reader := bufio.NewReader(os.Stdin)
	for {
		// 发送请求，调用 MyTest 接口
		response, err := grpcClient.SayHello(context.Background(), req)
		if err != nil {
			log.Fatal("发送请求失败，原因是:", err)
		}
		log.Println(response)

		reader.ReadLine()
	}
}
