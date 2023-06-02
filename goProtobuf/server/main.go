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

func main() {
	l, err := net.Listen("tcp", config.ServerPort)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	svc := grpc.NewServer()
	pb.RegisterAddressBookServiceServer(svc, new(addressBookService))
	slog.Info("start server")
	if err := svc.Serve(l); err != nil {
		slog.Error(err.Error())
		return
	}
}
