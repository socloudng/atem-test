package service

import (
	"atem-test/pbs/hello"
	"context"

	"google.golang.org/grpc"
)

type helloService struct {
	hello.MsgHelloServiceServer
}

func NewHelloService() *helloService {
	return &helloService{}
}

func (h *helloService) SayHello(ctx context.Context, req *hello.RequestMsg) (*hello.ResponseMsg, error) {
	return &hello.ResponseMsg{Message: "Hello " + req.Name + " , nice to meet you!"}, nil
}

func (h *helloService) RegisterHelloServer(serv *grpc.Server) {
	hello.RegisterMsgHelloServiceServer(serv, h)
}
