package service

import (
	"context"

	"golang.org/x/exp/slog"

	pb "github.com/emahiro/il/protobuf/pb/proto"
)

type AddressBookService struct{}

func (s *AddressBookService) GetPerson(ctx context.Context, in *pb.Person) (*pb.Person, error) {
	slog.InfoCtx(ctx, "GetPerson だよ")
	return &pb.Person{
		Name:  "Taro",
		Email: "taro@example.com",
	}, nil
}

func (s *AddressBookService) AddPerson(ctx context.Context, in *pb.Person) (*pb.Person, error) {
	slog.InfoCtx(ctx, "AddPerson だよ")
	return nil, nil
}
