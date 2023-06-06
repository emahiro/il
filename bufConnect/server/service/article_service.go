package service

import (
	"context"

	connect_go "github.com/bufbuild/connect-go"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc/metadata"

	v1 "github.com/emahiro/il/bufconnect/gen/proto/article/v1"
)

type ArticleService struct{}

func (s *ArticleService) GetArticle(ctx context.Context, req *connect_go.Request[v1.GetArticleRequest]) (*connect_go.Response[v1.GetArticleResponse], error) {
	slog.InfoCtx(ctx, "GetArticle Service invoked")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		slog.ErrorCtx(ctx, "metadata not found")
	}
	slog.InfoCtx(ctx, "metadata", md)

	header := req.Header()
	slog.InfoCtx(ctx, "request header", header)
	return connect_go.NewResponse(&v1.GetArticleResponse{
		Self: &v1.Article{
			Id:    1,
			Title: "title",
			Body:  "body",
		},
	}), nil
}

func (s *ArticleService) GetArticles(ctx context.Context, req *connect_go.Request[v1.GetArticlesRequest]) (*connect_go.Response[v1.GetArticlesResponse], error) {
	slog.InfoCtx(ctx, "GetArticles Service invoked")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		slog.ErrorCtx(ctx, "metadata not found")
	}
	slog.InfoCtx(ctx, "metadata", md)

	return connect_go.NewResponse(&v1.GetArticlesResponse{
		Lists: []*v1.Article{
			{
				Id:    1,
				Title: "title",
				Body:  "body",
			},
			{
				Id:    2,
				Title: "title2",
				Body:  "body2",
			},
		},
	}), nil
}
