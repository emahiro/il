package service

import (
	"context"
	"fmt"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc/metadata"

	pb "github.com/emahiro/il/protobuf/pb/proto"
)

type UserService struct{}

func (s *UserService) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	slog.InfoCtx(ctx, "GetUser だよ")
	return &pb.GetUserResponse{
		Self: &pb.User{
			Name: "Taro",
			Age:  20,
		},
	}, nil
}

func (s *UserService) GetUsers(ctx context.Context, in *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	slog.InfoCtx(ctx, "GetUsers だよ")
	metadata, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		slog.ErrorCtx(ctx, "failed to get metadata")
	}
	slog.InfoCtx(ctx, fmt.Sprintf("metadata: %v", metadata))
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
