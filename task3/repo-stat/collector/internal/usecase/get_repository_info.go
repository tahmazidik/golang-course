package usecase

import (
	"context"
	"errors"

	"repo-stat/collector/internal/domain"
)

type GetRepositoryInfo struct {
	repositoryInfoProvider RepositoryInfoProvider
}

func NewGetRepositoryInfo(repositoryInfoProvider RepositoryInfoProvider) *GetRepositoryInfo {
	return &GetRepositoryInfo{
		repositoryInfoProvider: repositoryInfoProvider,
	}
}

func (u *GetRepositoryInfo) Execute(ctx context.Context, url string) (*domain.RepositoryInfo, error) {
	if url == "" {
		return nil, errors.New("repository URL cannot be empty")
	}
	return u.repositoryInfoProvider.GetRepositoryInfo(ctx, url)
}
