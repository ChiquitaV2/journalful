package health

import (
	"context"
	healthpb "google.golang.org/grpc/health/grpc_health_v1" // Standard gRPC health package
	"sync"
)

type HealthCheckService struct {
	healthpb.UnimplementedHealthServer // Required for forward compatibility
	mu                                 sync.RWMutex
	// statusMap stores the serving status of different services.
	// Key is the service name, value is its ServingStatus.
	statusMap map[string]healthpb.HealthCheckResponse_ServingStatus
}

// NewHealthCheckService creates a new HealthCheckService instance.
func NewHealthCheckService() *HealthCheckService {
	return &HealthCheckService{
		statusMap: make(map[string]healthpb.HealthCheckResponse_ServingStatus),
	}
}

// SetServiceStatus allows changing the health status of a specific service.
// A common use case is to set the overall service status.
func (s *HealthCheckService) SetServiceStatus(service string, status healthpb.HealthCheckResponse_ServingStatus) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.statusMap[service] = status
}

// Check implements the HealthServer.Check RPC method.
// It returns the current serving status of a service.
func (s *HealthCheckService) Check(ctx context.Context, in *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// If no specific service is requested, check the overall service health.
	// You might have a default service name (e.g., "").
	serviceName := in.GetService()
	if serviceName == "" {
		serviceName = "" // Represents the overall service health
	}

	status, ok := s.statusMap[serviceName]
	if !ok {
		// If the service is not found or not registered, consider it NOT_SERVING or UNKNOWN
		return &healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_NOT_SERVING}, nil
	}

	return &healthpb.HealthCheckResponse{Status: status}, nil
}

func (s *HealthCheckService) mustEmbedUnimplementedHealthServer() {
	// This method is required to ensure that the server implements the
	// UnimplementedHealthServer interface for forward compatibility.
	// It can be left empty, as the UnimplementedHealthServer provides default
	// implementations for all methods.
}
