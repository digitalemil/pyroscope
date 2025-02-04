// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: util/httpgrpc/httpgrpc.proto

package httpgrpcconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	httpgrpc "github.com/grafana/pyroscope/pkg/util/httpgrpc"
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
	// HTTPName is the fully-qualified name of the HTTP service.
	HTTPName = "httpgrpc.HTTP"
)

// HTTPClient is a client for the httpgrpc.HTTP service.
type HTTPClient interface {
	Handle(context.Context, *connect_go.Request[httpgrpc.HTTPRequest]) (*connect_go.Response[httpgrpc.HTTPResponse], error)
}

// NewHTTPClient constructs a client for the httpgrpc.HTTP service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewHTTPClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) HTTPClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &hTTPClient{
		handle: connect_go.NewClient[httpgrpc.HTTPRequest, httpgrpc.HTTPResponse](
			httpClient,
			baseURL+"/httpgrpc.HTTP/Handle",
			opts...,
		),
	}
}

// hTTPClient implements HTTPClient.
type hTTPClient struct {
	handle *connect_go.Client[httpgrpc.HTTPRequest, httpgrpc.HTTPResponse]
}

// Handle calls httpgrpc.HTTP.Handle.
func (c *hTTPClient) Handle(ctx context.Context, req *connect_go.Request[httpgrpc.HTTPRequest]) (*connect_go.Response[httpgrpc.HTTPResponse], error) {
	return c.handle.CallUnary(ctx, req)
}

// HTTPHandler is an implementation of the httpgrpc.HTTP service.
type HTTPHandler interface {
	Handle(context.Context, *connect_go.Request[httpgrpc.HTTPRequest]) (*connect_go.Response[httpgrpc.HTTPResponse], error)
}

// NewHTTPHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewHTTPHandler(svc HTTPHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/httpgrpc.HTTP/Handle", connect_go.NewUnaryHandler(
		"/httpgrpc.HTTP/Handle",
		svc.Handle,
		opts...,
	))
	return "/httpgrpc.HTTP/", mux
}

// UnimplementedHTTPHandler returns CodeUnimplemented from all methods.
type UnimplementedHTTPHandler struct{}

func (UnimplementedHTTPHandler) Handle(context.Context, *connect_go.Request[httpgrpc.HTTPRequest]) (*connect_go.Response[httpgrpc.HTTPResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("httpgrpc.HTTP.Handle is not implemented"))
}
