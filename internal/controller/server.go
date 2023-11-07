package controller

import (
	"LinkEnshorter/pb"
	"context"
	"database/sql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedUrlShorterServer
	service service
}

func NewServer(service service) *Server {
	return &Server{service: service}
}

func (s *Server) SaveURL(ctx context.Context, in *pb.Url) (*pb.Hash, error) {
	hash, err := s.service.SaveShortURL(ctx, in.Url)
	if err != nil {
		return &pb.Hash{}, status.Error(codes.Internal, err.Error())
	}
	return &pb.Hash{Hash: hash}, nil
}

func (s *Server) ShowURL(ctx context.Context, in *pb.Hash) (*pb.Url, error) {
	url, err := s.service.ShowLink(ctx, in.Hash)
	if err == sql.ErrNoRows || err != nil && err.Error() == "no such hash in cache" {
		return &pb.Url{}, status.Error(codes.InvalidArgument, err.Error())
	} else if err != nil {
		return &pb.Url{}, status.Error(codes.Internal, err.Error())
	}
	return &pb.Url{Url: url}, nil
}
