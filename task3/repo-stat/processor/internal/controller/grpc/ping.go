package grpc

import (
	"context"
	"log/slog"

	"repo-stat/processor/internal/usecase"
	processorpb "repo-stat/proto/processor"
)

type Server struct {
	processorpb.UnimplementedProcessorServer
	log               *slog.Logger
	ping              *usecase.Ping
	getRepositoryInfo *usecase.GetRepositoryInfo
}

func NewServer(log *slog.Logger, ping *usecase.Ping, getRepositoryInfo *usecase.GetRepositoryInfo) *Server {
	return &Server{
		log:               log,
		ping:              ping,
		getRepositoryInfo: getRepositoryInfo,
	}
}

func (s *Server) Ping(ctx context.Context, _ *processorpb.PingRequest) (*processorpb.PingResponse, error) {
	s.log.Debug("processorpb ping request received")
	return &processorpb.PingResponse{
		Reply: s.ping.Execute(ctx),
	}, nil
}
