package usecase

import (
	"context"
	"repo-stat/api/internal/domain"
)

type SubscriberPinger interface {
	Ping(ctx context.Context) domain.PingStatus
}

type ProcessorPinger interface {
	Ping(ctx context.Context) domain.PingStatus
}

type ProcessorRepositoryInfoGetter interface {
	GetRepositoryInfo(ctx context.Context, url string) (*domain.RepositoryInfo, error)
}
