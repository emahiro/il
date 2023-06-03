package main

import (
	"context"
	"net"

	"github.com/emahiro/il/protobuf/config"
	pb "github.com/emahiro/il/protobuf/pb/proto"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

type addressBookService struct{}

func (s *addressBookService) GetPerson(ctx context.Context, in *pb.Person) (*pb.Person, error) {
	slog.InfoCtx(ctx, "GetPerson だよ")
	return &pb.Person{
		Name:  "Taro",
		Email: "taro@example.com",
	}, nil
}

func (s *addressBookService) AddPerson(ctx context.Context, in *pb.Person) (*pb.Person, error) {
	slog.InfoCtx(ctx, "AddPerson だよ")
	return nil, nil
}

type userService struct{}

func (s *userService) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	slog.InfoCtx(ctx, "GetUser だよ")
	return &pb.GetUserResponse{
		Self: &pb.User{
			Name: "Taro",
			Age:  20,
		},
	}, nil
}

func (s *userService) GetUsers(ctx context.Context, in *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	slog.InfoCtx(ctx, "GetUsers だよ")
	return &pb.GetUsersResponse{
		Lists: []*pb.User{
			{
				Name: "Taro",
				Age:  20,
			},
			{
				Name: "Jiro",
				Age:  30,
			},
		},
	}, nil
}

func main() {
	l, err := net.Listen("tcp", config.ServerPort)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	svc := grpc.NewServer()
	pb.RegisterAddressBookServiceServer(svc, new(addressBookService))
	pb.RegisterUserServiceServer(svc, new(userService))
	slog.Info("start server")
	if err := svc.Serve(l); err != nil {
		slog.Error(err.Error())
		return
	}
}
