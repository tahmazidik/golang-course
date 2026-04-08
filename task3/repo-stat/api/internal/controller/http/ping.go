package http

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"repo-stat/api/internal/domain"
	"repo-stat/api/internal/dto"
	"repo-stat/api/internal/usecase"
)

func NewPingHandler(log *slog.Logger, ping *usecase.Ping) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result := ping.Execute(r.Context())

		response := dto.PingResponse{
			Status: string(result.Status),
			Services: func() []dto.ServiceStatusResponse {
				service := make([]dto.ServiceStatusResponse, len(result.Services))
				for i, s := range result.Services {
					service[i] = dto.ServiceStatusResponse{
						Name:   s.Name,
						Status: string(s.Status),
					}
				}
				return service
			}(),
		}

		w.Header().Set("Content-Type", "application/json")
		statusCode := http.StatusOK
		if result.Status == domain.OverallStatusDegraded {
			statusCode = http.StatusServiceUnavailable
		}
		w.WriteHeader(statusCode)

		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Error("failed to write ping response", "error", err)
		}
	}
}
