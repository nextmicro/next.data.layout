package service

import (
	"context"

	pb "next.data.layout/api/shorturl/v1"
)

type ShortUrlService struct {
	pb.UnimplementedShortUrlServer
}

func NewShortUrlService() *ShortUrlService {
	return &ShortUrlService{}
}

func (s *ShortUrlService) Expand(ctx context.Context, req *pb.ExpandRequest) (*pb.ExpandReply, error) {
	return &pb.ExpandReply{
		Url: req.Shorten,
	}, nil
}
func (s *ShortUrlService) Shorten(ctx context.Context, req *pb.ShortenRequest) (*pb.ShortenReply, error) {
	return &pb.ShortenReply{
		Shorten: req.Url,
	}, nil
}
