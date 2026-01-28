package service

import (
	"context"
	"fmt"
	"srtk/proto/accounting/userpb"
)

type AccountingServer struct {
	userpb.UnimplementedUserServiceServer
	DB map[int64]*userpb.CreateUserRequest
}

func NewAccountingServer() *AccountingServer {
	return &AccountingServer{
		DB: make(map[int64]*userpb.CreateUserRequest),
	}
}

// gRPC метод
func (s *AccountingServer) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	s.DB[req.Id] = req
	fmt.Println("User saved in inventory-service:", req)
	return &userpb.CreateUserResponse{Ok: true}, nil
}
