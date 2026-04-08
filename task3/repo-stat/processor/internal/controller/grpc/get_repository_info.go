package grpc

import (
	"context"
	processorpb "repo-stat/proto/processor"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) GetRepositoryInfo(ctx context.Context, req *processorpb.GetRepositoryInfoRequest) (*processorpb.GetRepositoryInfoResponse, error) {
	s.log.Debug("processor get repository info request received", "url", req.Url)
	info, err := s.getRepositoryInfo.Execute(ctx, req.Url)
	if err != nil {
		s.log.Error("failed to get repository info", "error", err)
		return nil, err
	}
	return &processorpb.GetRepositoryInfoResponse{
		FullName:    info.FullName,
		Description: info.Description,
		Stars:       info.Stars,
		Forks:       info.Forks,
		CreatedAt:   timestamppb.New(info.CreatedAt),
	}, nil
}
