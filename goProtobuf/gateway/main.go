package main

import (
	"context"
	"flag"
	"net/http"

	"github.com/emahiro/il/protobuf/config"
	gw "github.com/emahiro/il/protobuf/pb/proto"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpcServerEp = flag.String("grpc-server-endpoint", "localhost"+config.ServerPort, "gRPC server endpoint")
)

// gateway
func run(ctx context.Context) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gw.RegisterAddressBookServiceHandlerFromEndpoint(ctx, mux, *grpcServerEp, opts)
	if err != nil {
		return err
	}
	glog.Infof("[INFO]start server...")
	return http.ListenAndServe(config.GatewayPort, mux)
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
