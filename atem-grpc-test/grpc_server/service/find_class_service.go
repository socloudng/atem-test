package service

import (
	"atem-test/pbs/student"
	"context"

	"google.golang.org/grpc"
)

type findClassService struct {
	student.FindClassServer
}

func NewFindClassService() *findClassService {
	return &findClassService{}
}

func (h *findClassService) MyClass(ctx context.Context, req *student.Student) (*student.Class, error) {
	return &student.Class{ClassName: "就不告诉你" + req.Name}, nil
}

func (h *findClassService) RegisterFindClassServer(serv *grpc.Server) {
	student.RegisterFindClassServer(serv, h)
}
