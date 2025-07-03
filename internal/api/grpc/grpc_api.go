package grpc

import (
	"database/sql"
	"github.com/chiquitav2/journalful/internal/api"
	articleImp "github.com/chiquitav2/journalful/internal/article"
	"github.com/chiquitav2/journalful/internal/health"
	libraryImp "github.com/chiquitav2/journalful/internal/library"
	profileImp "github.com/chiquitav2/journalful/internal/profile"
	article "github.com/chiquitav2/journalful/pkg/articles/v1"
	"github.com/chiquitav2/journalful/pkg/library/v1"
	"github.com/chiquitav2/journalful/pkg/profile/v1"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1" // Standard gRPC health package
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
)

type GrpcService struct {
	api.ApiModule
	dbConn *sql.DB
	server *grpc.Server
}

func NewGrpcService(conn *sql.DB) *GrpcService {
	return &GrpcService{
		dbConn: conn,
	}
}

func (g *GrpcService) Register() error {
	g.server = grpc.NewServer()
	//enable grpc reflection for debugging
	reflection.Register(g.server)

	// Register the article service with the gRPC server
	article.RegisterArticlesServiceServer(g.server, articleImp.NewArticleGrpcService(g.dbConn))
	profile.RegisterAuthorServiceServer(g.server, profileImp.NewAuthorService(g.dbConn))
	profile.RegisterProfileServiceServer(g.server, profileImp.NewProfileService(g.dbConn))
	library.RegisterLibraryServiceServer(g.server, libraryImp.NewLibraryService(g.dbConn))

	healthSvr := health.NewHealthCheckService()
	healthSvr.SetServiceStatus("", healthpb.HealthCheckResponse_SERVING)         // Set overall service status to SERVING
	healthSvr.SetServiceStatus("articles", healthpb.HealthCheckResponse_SERVING) // Set articles service status to SERVING
	healthpb.RegisterHealthServer(g.server, healthSvr)

	return nil
}

func (g *GrpcService) Start() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		slog.Error("Failed to listen: %v", err)
	}
	slog.Info("gRPC server is starting on port 50051")
	if err := g.server.Serve(lis); err != nil {
		slog.Error("Failed to serve: %v", err)
		return err
	}
	return nil
}
