package usecase

import (
	"context"
	"repo-stat/processor/internal/domain"
)

type GetRepositoryInfo struct {
	collector Collector
}

func NewGetRepositoryInfo(collecot Collector) *GetRepositoryInfo {
	return &GetRepositoryInfo{
		collector: collecot,
	}
}

func (g *GetRepositoryInfo) Execute(ctx context.Context, url string) (*domain.RepositoryInfo, error) {
	return g.collector.GetRepositoryInfo(ctx, url)
}
