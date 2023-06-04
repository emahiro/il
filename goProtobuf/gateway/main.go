package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/emahiro/il/protobuf/config"
	gw "github.com/emahiro/il/protobuf/pb/proto"
)

func newGateway(ctx context.Context, conn *grpc.ClientConn, opts ...runtime.ServeMuxOption) (http.Handler, error) {
	mux := runtime.NewServeMux(opts...)

	for _, fn := range []func(context.Context, *runtime.ServeMux, *grpc.ClientConn) error{
		gw.RegisterAddressBookServiceHandler,
		gw.RegisterUserServiceHandler,
	} {
		if err := fn(ctx, mux, conn); err != nil {
			return nil, err
		}
	}

	return mux, nil
}

func dial(ctx context.Context, addr string) (*grpc.ClientConn, error) {
	return grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	conn, err := dial(ctx, config.ServerPort)
	if err != nil {
		slog.ErrorCtx(ctx, "failed to open connecton. err: %v", err)
		return
	}
	defer func() {
		if err := conn.Close(); err != nil {
			slog.ErrorCtx(ctx, "failed to close connecton. err: %v", err)
		}
	}()

	mux := http.NewServeMux()
	gw, err := newGateway(ctx, conn)
	if err != nil {
		slog.ErrorCtx(ctx, "failed to create gateway. err: %v", err)
		return
	}
	mux.Handle("/", gw)

	gwServer := &http.Server{
		Addr:    config.GatewayPort,
		Handler: mux,
	}

	slog.InfoCtx(ctx, fmt.Sprintf("start gateway server at localhost%s", config.GatewayPort))
	if err := gwServer.ListenAndServe(); err != nil {
		slog.ErrorCtx(ctx, "failed to start gateway server. err: %v", err)
		return
	}
	defer func() {
		if err := gwServer.Shutdown(ctx); err != nil {
			slog.ErrorCtx(ctx, "failed to shutdown gateway server. err: %v", err)
		}
	}()
}
