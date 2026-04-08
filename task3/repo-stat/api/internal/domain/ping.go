package domain

type PingStatus string

type OverallStatus string

const (
	PingStatusUp          PingStatus    = "up"
	PingStatusDown        PingStatus    = "down"
	OverallStatusOk       OverallStatus = "ok"
	OverallStatusDegraded OverallStatus = "degraded"
)

type ServiceStatus struct {
	Name   string     `json:"name"`
	Status PingStatus `json:"status"`
}
type PingResult struct {
	Status   OverallStatus   `json:"status"`
	Services []ServiceStatus `json:"services"`
}
