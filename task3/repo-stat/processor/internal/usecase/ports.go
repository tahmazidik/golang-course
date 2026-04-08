package usecase

import (
	"context"
	"repo-stat/processor/internal/domain"
)

type Collector interface {
	GetRepositoryInfo(ctx context.Context, url string) (*domain.RepositoryInfo, error)
}
