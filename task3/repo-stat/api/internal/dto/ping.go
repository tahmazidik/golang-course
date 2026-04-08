package dto

import "time"

type PingResponse struct {
	Status   string                  `json:"status"`
	Services []ServiceStatusResponse `json:"services"`
}

type ServiceStatusResponse struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type RepositoryInfoResponse struct {
	FullName    string    `json:"full_name"`
	Description string    `json:"description"`
	Stars       int64     `json:"stars"`
	Forks       int64     `json:"forks"`
	CreatedAt   time.Time `json:"created_at"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
