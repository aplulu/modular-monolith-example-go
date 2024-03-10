// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: example/article/v1/article.proto

package articleconnect

import (
	context "context"
	errors "errors"
	v1 "github.com/aplulu/modular-monolith-example-go/internal/grpc/example/article/v1"
	connect_go "github.com/bufbuild/connect-go"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion1_7_0

const (
	// ArticleServiceName is the fully-qualified name of the ArticleService service.
	ArticleServiceName = "example.article.v1.ArticleService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ArticleServiceListArticleProcedure is the fully-qualified name of the ArticleService's
	// ListArticle RPC.
	ArticleServiceListArticleProcedure = "/example.article.v1.ArticleService/ListArticle"
)

// ArticleServiceClient is a client for the example.article.v1.ArticleService service.
type ArticleServiceClient interface {
	// ユーザを取得
	ListArticle(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.ListArticleResponse], error)
}

// NewArticleServiceClient constructs a client for the example.article.v1.ArticleService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewArticleServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ArticleServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &articleServiceClient{
		listArticle: connect_go.NewClient[emptypb.Empty, v1.ListArticleResponse](
			httpClient,
			baseURL+ArticleServiceListArticleProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
	}
}

// articleServiceClient implements ArticleServiceClient.
type articleServiceClient struct {
	listArticle *connect_go.Client[emptypb.Empty, v1.ListArticleResponse]
}

// ListArticle calls example.article.v1.ArticleService.ListArticle.
func (c *articleServiceClient) ListArticle(ctx context.Context, req *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.ListArticleResponse], error) {
	return c.listArticle.CallUnary(ctx, req)
}

// ArticleServiceHandler is an implementation of the example.article.v1.ArticleService service.
type ArticleServiceHandler interface {
	// ユーザを取得
	ListArticle(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.ListArticleResponse], error)
}

// NewArticleServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewArticleServiceHandler(svc ArticleServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	articleServiceListArticleHandler := connect_go.NewUnaryHandler(
		ArticleServiceListArticleProcedure,
		svc.ListArticle,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	)
	return "/example.article.v1.ArticleService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ArticleServiceListArticleProcedure:
			articleServiceListArticleHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedArticleServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedArticleServiceHandler struct{}

func (UnimplementedArticleServiceHandler) ListArticle(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.ListArticleResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("example.article.v1.ArticleService.ListArticle is not implemented"))
}
