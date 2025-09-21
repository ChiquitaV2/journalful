package grpcapi // Renamed package for clarity and to avoid conflicts.

import (
	"context"
	"database/sql"
	"fmt" // For error wrapping
	"log/slog"
	"net"

	"github.com/chiquitav2/journalful/internal/auth"
	libraryImp "github.com/chiquitav2/journalful/internal/library"
	profileImp "github.com/chiquitav2/journalful/internal/profile"
	"github.com/chiquitav2/journalful/pkg/library/v1"
	"github.com/chiquitav2/journalful/pkg/profile/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	dbConn   *sql.DB
	server   *grpc.Server
	health   *health.Server
	config   *conf.Config
	certFile string
	keyFile  string
}

// NewServer creates a new gRPC server.
// Renamed from NewGrpcService.
func NewServer(conn *sql.DB, cfg *conf.Config, certFile, keyFile string) *Server {
	return &Server{
		dbConn:   conn,
		health:   health.NewServer(), // Initialize health server here.
		config:   cfg,
		certFile: certFile,
		keyFile:  keyFile,
	}
}

// Register registers all the gRPC services and reflection.
func (s *Server) Register() error {
	authorizer, err := auth.NewZitadelAuthorizer(context.Background(), s.config.Zitadel.Domain, s.config.Zitadel.KeyPath)
	if err != nil {
		return fmt.Errorf("failed to create zitadel authorizer: %w", err)
	}

	authInterceptor := NewAuthInterceptor(*authorizer)

	creds, err := credentials.NewServerTLSFromFile(s.certFile, s.keyFile)
	if err != nil {
		return fmt.Errorf("failed to load TLS keys: %w", err)
	}

	s.server = grpc.NewServer(
		grpc.Creds(creds),
		grpc.ChainUnaryInterceptor(
			RecoveryInterceptor,
			LoggingInterceptor,
			ErrorInterceptor,
			authInterceptor.Unary(),
		),
	)
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
		return fmt.Errorf("failed to listen on %s: %w", addr, err)
	}

	slog.Info("gRPC server starting", "address", lis.Addr().String())
	if err := s.server.Serve(lis); err != nil {
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
	slog.Info("gRPC server stopped")
}
