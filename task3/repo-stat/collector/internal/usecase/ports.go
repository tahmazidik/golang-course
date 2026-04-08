package usecase

import (
	"context"
	"repo-stat/collector/internal/domain"
)

type RepositoryInfoProvider interface {
	GetRepositoryInfo(ctx context.Context, url string) (*domain.RepositoryInfo, error)
}
