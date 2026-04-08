package usecase

import (
	"context"
	"repo-stat/api/internal/domain"
)

type Ping struct {
	processorPinger  ProcessorPinger
	subscriberPinger SubscriberPinger
}

func NewPing(processorPinger ProcessorPinger, subscriberPinger SubscriberPinger) *Ping {
	return &Ping{
		processorPinger:  processorPinger,
		subscriberPinger: subscriberPinger,
	}
}

func (u *Ping) Execute(ctx context.Context) domain.PingResult {
	processorStatus := u.processorPinger.Ping(ctx)
	subscriberStatus := u.subscriberPinger.Ping(ctx)

	services := []domain.ServiceStatus{
		{
			Name:   "processor",
			Status: processorStatus,
		},
		{
			Name:   "subscriber",
			Status: subscriberStatus,
		},
	}

	overallStatus := domain.OverallStatusOk

	for _, service := range services {
		if service.Status == domain.PingStatusDown {
			overallStatus = domain.OverallStatusDegraded
			break
		}
	}

	return domain.PingResult{
		Status:   overallStatus,
		Services: services,
	}
}
