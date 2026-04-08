package grpc

import (
	"context"
	"log/slog"
	"repo-stat/collector/internal/usecase"
	collectorpb "repo-stat/proto/collector"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	collectorpb.UnimplementedCollectorServer
	log               *slog.Logger
	getRepositoryInfo *usecase.GetRepositoryInfo
}

func NewServer(log *slog.Logger, getRepositoryInfo *usecase.GetRepositoryInfo) *Server {
	return &Server{
		log:               log,
		getRepositoryInfo: getRepositoryInfo,
	}
}

func (s *Server) GetRepositoryInfo(ctx context.Context, req *collectorpb.GetRepositoryInfoRequest) (*collectorpb.GetRepositoryInfoResponse, error) {
	s.log.Debug("collector get repository info request received", "url", req.Url)
	info, err := s.getRepositoryInfo.Execute(ctx, req.Url)
	if err != nil {
		s.log.Error("failed to get repository info", "error", err)
		return nil, err
	}
	return &collectorpb.GetRepositoryInfoResponse{
		FullName:    info.FullName,
		Description: info.Description,
		Stars:       info.Stars,
		Forks:       info.Forks,
		CreatedAt:   timestamppb.New(info.CreatedAt),
	}, nil
}
