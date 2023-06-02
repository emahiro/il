package main

import (
	"context"
	"flag"
	"net/http"

	pb "github.com/emahiro/il/protobuf/pb/proto"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpcServerEp = flag.String("grpc-server-endpoint", "localhost:50051", "gRPC server endpoint")
)

func run(ctx context.Context) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterAddressBookServiceHandlerFromEndpoint(ctx, mux, *grpcServerEp, opts)
	if err != nil {
		return err
	}
	return http.ListenAndServe(":8082", mux)
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	flag.Parse()
	defer glog.Flush()

	if err := run(ctx); err != nil {
		glog.Fatal(err)
	}
}
