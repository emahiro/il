package main

import (
	"context"
	"net"

	"github.com/golang/glog"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"

	"github.com/emahiro/il/protobuf/config"
	pb "github.com/emahiro/il/protobuf/pb/proto"
	"github.com/emahiro/il/protobuf/server/service"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	l, err := net.Listen("tcp", config.ServerPort)
	defer func() {
		if err != nil {
			glog.Errorf("failed to close server. err: %v", err)
		}
	}()

	svr := grpc.NewServer()
	pb.RegisterAddressBookServiceServer(svr, new(service.AddressBookService))
	pb.RegisterUserServiceServer(svr, new(service.UserService))
	slog.Info("start server")

	go func() {
		defer svr.GracefulStop()
		<-ctx.Done()
	}()

	if err := svr.Serve(l); err != nil {
		slog.Error(err.Error())
		return
	}
}
