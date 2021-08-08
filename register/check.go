package register

import (
	"context"

	"layout/pkg/log"

	"google.golang.org/grpc/health/grpc_health_v1"
)

// Healthy check register health.
type Healthy struct {
	Status grpc_health_v1.HealthCheckResponse_ServingStatus
	Reason string
}

// Watch .
func (h *Healthy) Watch(*grpc_health_v1.HealthCheckRequest, grpc_health_v1.Health_WatchServer) error {
	return nil
}

// OffLine .
func (h *Healthy) OffLine(reason string) {
	h.Status = grpc_health_v1.HealthCheckResponse_NOT_SERVING
	h.Reason = reason
	log.Errorf(reason)
}

// OnLine .
func (h *Healthy) OnLine(reason string) {
	h.Status = grpc_health_v1.HealthCheckResponse_SERVING
	h.Reason = reason
	log.Errorf(reason)
}

// Check .
func (h *Healthy) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: h.Status,
	}, nil
}
