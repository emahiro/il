// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/article/v1/article.proto

package genconnect

import (
	gen "/gen"
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// ArticleServiceName is the fully-qualified name of the ArticleService service.
	ArticleServiceName = "article.v1.ArticleService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ArticleServiceGetArticleProcedure is the fully-qualified name of the ArticleService's GetArticle
	// RPC.
	ArticleServiceGetArticleProcedure = "/article.v1.ArticleService/GetArticle"
	// ArticleServiceGetArticlesProcedure is the fully-qualified name of the ArticleService's
	// GetArticles RPC.
	ArticleServiceGetArticlesProcedure = "/article.v1.ArticleService/GetArticles"
)

// ArticleServiceClient is a client for the article.v1.ArticleService service.
type ArticleServiceClient interface {
	GetArticle(context.Context, *connect_go.Request[gen.GetArticleRequest]) (*connect_go.Response[gen.GetArticleResponse], error)
	GetArticles(context.Context, *connect_go.Request[gen.GetArticlesRequest]) (*connect_go.ServerStreamForClient[gen.GetArticlesResponse], error)
}

// NewArticleServiceClient constructs a client for the article.v1.ArticleService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewArticleServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ArticleServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &articleServiceClient{
		getArticle: connect_go.NewClient[gen.GetArticleRequest, gen.GetArticleResponse](
			httpClient,
			baseURL+ArticleServiceGetArticleProcedure,
			opts...,
		),
		getArticles: connect_go.NewClient[gen.GetArticlesRequest, gen.GetArticlesResponse](
			httpClient,
			baseURL+ArticleServiceGetArticlesProcedure,
			opts...,
		),
	}
}

// articleServiceClient implements ArticleServiceClient.
type articleServiceClient struct {
	getArticle  *connect_go.Client[gen.GetArticleRequest, gen.GetArticleResponse]
	getArticles *connect_go.Client[gen.GetArticlesRequest, gen.GetArticlesResponse]
}

// GetArticle calls article.v1.ArticleService.GetArticle.
func (c *articleServiceClient) GetArticle(ctx context.Context, req *connect_go.Request[gen.GetArticleRequest]) (*connect_go.Response[gen.GetArticleResponse], error) {
	return c.getArticle.CallUnary(ctx, req)
}

// GetArticles calls article.v1.ArticleService.GetArticles.
func (c *articleServiceClient) GetArticles(ctx context.Context, req *connect_go.Request[gen.GetArticlesRequest]) (*connect_go.ServerStreamForClient[gen.GetArticlesResponse], error) {
	return c.getArticles.CallServerStream(ctx, req)
}

// ArticleServiceHandler is an implementation of the article.v1.ArticleService service.
type ArticleServiceHandler interface {
	GetArticle(context.Context, *connect_go.Request[gen.GetArticleRequest]) (*connect_go.Response[gen.GetArticleResponse], error)
	GetArticles(context.Context, *connect_go.Request[gen.GetArticlesRequest], *connect_go.ServerStream[gen.GetArticlesResponse]) error
}

// NewArticleServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewArticleServiceHandler(svc ArticleServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(ArticleServiceGetArticleProcedure, connect_go.NewUnaryHandler(
		ArticleServiceGetArticleProcedure,
		svc.GetArticle,
		opts...,
	))
	mux.Handle(ArticleServiceGetArticlesProcedure, connect_go.NewServerStreamHandler(
		ArticleServiceGetArticlesProcedure,
		svc.GetArticles,
		opts...,
	))
	return "/article.v1.ArticleService/", mux
}

// UnimplementedArticleServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedArticleServiceHandler struct{}

func (UnimplementedArticleServiceHandler) GetArticle(context.Context, *connect_go.Request[gen.GetArticleRequest]) (*connect_go.Response[gen.GetArticleResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("article.v1.ArticleService.GetArticle is not implemented"))
}

func (UnimplementedArticleServiceHandler) GetArticles(context.Context, *connect_go.Request[gen.GetArticlesRequest], *connect_go.ServerStream[gen.GetArticlesResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("article.v1.ArticleService.GetArticles is not implemented"))
}
