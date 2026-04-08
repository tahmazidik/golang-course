package usecase

import (
	"context"
	"repo-stat/api/internal/domain"
)

type GetRepositoryInfo struct {
	processore ProcessorRepositoryInfoGetter
}

func NewGetRepositoryInfo(processor ProcessorRepositoryInfoGetter) *GetRepositoryInfo {
	return &GetRepositoryInfo{
		processore: processor,
	}
}

func (g *GetRepositoryInfo) Execute(ctx context.Context, url string) (*domain.RepositoryInfo, error) {
	return g.processore.GetRepositoryInfo(ctx, url)
}
