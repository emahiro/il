package main

import (
	"context"
	"net/http"

	"golang.org/x/exp/slog"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/emahiro/il/bufconnect/gen/proto/article/v1/articlev1connect"
	"github.com/emahiro/il/bufconnect/server/service"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	articler := &service.ArticleService{}

	mux := http.NewServeMux()
	path, handler := articlev1connect.NewArticleServiceHandler(articler)
	mux.Handle(path, handler)

	server := &http.Server{
		Addr: ":8080",
		// use http2 without TLS
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}

	go func() {
		<-ctx.Done()
		if err := server.Close(); err != nil {
			panic(err)
		}
	}()

	slog.InfoCtx(ctx, "[INFO] server started", "addr", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
