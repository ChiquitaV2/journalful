package grpcapi // Renamed package for clarity and to avoid conflicts.

import (
	"database/sql"
	"fmt" // For error wrapping
	libraryImp "github.com/chiquitav2/journalful/internal/library"
	profileImp "github.com/chiquitav2/journalful/internal/profile"
	"github.com/chiquitav2/journalful/pkg/library/v1"
	"github.com/chiquitav2/journalful/pkg/profile/v1"
	"log/slog"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	"github.com/chiquitav2/journalful/internal/api"
	articleImp "github.com/chiquitav2/journalful/internal/article"
	article "github.com/chiquitav2/journalful/pkg/articles/v1"
	"github.com/chiquitav2/journalful/pkg/conf"
)

// Server manages the gRPC server and its services.
// The name is changed from GrpcService to Server for brevity, as it's in a grpcapi package.
type Server struct {
	api.ApiModule
	dbConn *sql.DB
	server *grpc.Server
	health *health.Server
}

// NewServer creates a new gRPC server.
// Renamed from NewGrpcService.
func NewServer(conn *sql.DB) *Server {
	return &Server{
		dbConn: conn,
		health: health.NewServer(), // Initialize health server here.
	}
}

// Register registers all the gRPC services and reflection.
func (s *Server) Register() error {
	s.server = grpc.NewServer()
	// Enable gRPC reflection for debugging.
	reflection.Register(s.server)

	// Register services.
	article.RegisterArticlesServiceServer(s.server, articleImp.NewArticleGrpcHandler(s.dbConn))
	profile.RegisterAuthorServiceServer(s.server, profileImp.NewProfileGrpcHandler(s.dbConn))
	profile.RegisterProfileServiceServer(s.server, profileImp.NewProfileGrpcHandler(s.dbConn))
	library.RegisterLibraryServiceServer(s.server, libraryImp.NewLibraryGrpcHandler(s.dbConn))

	// Register health check service.
	healthpb.RegisterHealthServer(s.server, s.health)
	// Set initial status for services.
	s.health.SetServingStatus("articles.ArticlesService", healthpb.HealthCheckResponse_SERVING)
	s.health.SetServingStatus("profile.AuthorService", healthpb.HealthCheckResponse_SERVING)
	s.health.SetServingStatus("profile.ProfileService", healthpb.HealthCheckResponse_SERVING)
	s.health.SetServingStatus("library.LibraryService", healthpb.HealthCheckResponse_SERVING)
	s.health.SetServingStatus("", healthpb.HealthCheckResponse_SERVING) // Overall server status.

	return nil
}

// Start starts the gRPC server.
func (s *Server) Start(cfg *conf.Config) error {
	addr := fmt.Sprintf(":%s", cfg.Server.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		// Correct error handling: return the error instead of just logging.
		// Also using structured logging.
		slog.Error("failed to listen", "address", addr, "error", err)
		return fmt.Errorf("failed to listen on %s: %w", addr, err)
	}

	slog.Info("gRPC server starting", "address", lis.Addr().String())
	// This is a blocking call.
	if err := s.server.Serve(lis); err != nil {
		slog.Error("failed to serve gRPC", "error", err)
		return fmt.Errorf("failed to serve gRPC: %w", err)
	}

	return nil
}

// Stop gracefully stops the gRPC server.
func (s *Server) Stop() {
	slog.Info("stopping gRPC server")
	// Set all services to NOT_SERVING.
	s.health.Shutdown()
	s.server.GracefulStop()
}
