// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v1.0.0
// - protoc             v3.21.12
// source: api/shorturl/v1/shorturl.proto

package v1

import (
	context "context"
	http "github.com/nextmicro/next/transport/http"
	binding "github.com/nextmicro/next/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the nextmicro package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationShortUrlExpand = "/api.shorturl.v1.ShortUrl/Expand"
const OperationShortUrlShorten = "/api.shorturl.v1.ShortUrl/Shorten"

type ShortUrlHTTPServer interface {
	Expand(context.Context, *ExpandRequest) (*ExpandReply, error)
	Shorten(context.Context, *ShortenRequest) (*ShortenReply, error)
}

func RegisterShortUrlHTTPServer(s *http.Server, srv ShortUrlHTTPServer) {
	r := s.Route("/")
	r.GET("/expand/{shorten}", _ShortUrl_Expand0_HTTP_Handler(srv))
	r.POST("/shorten/{url}", _ShortUrl_Shorten0_HTTP_Handler(srv))
}

func _ShortUrl_Expand0_HTTP_Handler(srv ShortUrlHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ExpandRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShortUrlExpand)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Expand(ctx, req.(*ExpandRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ExpandReply)
		return ctx.Result(200, reply)
	}
}

func _ShortUrl_Shorten0_HTTP_Handler(srv ShortUrlHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ShortenRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShortUrlShorten)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Shorten(ctx, req.(*ShortenRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ShortenReply)
		return ctx.Result(200, reply)
	}
}

type ShortUrlHTTPClient interface {
	Expand(ctx context.Context, req *ExpandRequest, opts ...http.CallOption) (rsp *ExpandReply, err error)
	Shorten(ctx context.Context, req *ShortenRequest, opts ...http.CallOption) (rsp *ShortenReply, err error)
}

type ShortUrlHTTPClientImpl struct {
	cc *http.Client
}

func NewShortUrlHTTPClient(client *http.Client) ShortUrlHTTPClient {
	return &ShortUrlHTTPClientImpl{client}
}

func (c *ShortUrlHTTPClientImpl) Expand(ctx context.Context, in *ExpandRequest, opts ...http.CallOption) (*ExpandReply, error) {
	var out ExpandReply
	pattern := "/expand/{shorten}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationShortUrlExpand))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShortUrlHTTPClientImpl) Shorten(ctx context.Context, in *ShortenRequest, opts ...http.CallOption) (*ShortenReply, error) {
	var out ShortenReply
	pattern := "/shorten/{url}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationShortUrlShorten))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
