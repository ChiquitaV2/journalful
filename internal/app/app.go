package app

import (
	"database/sql"
	_ "embed"
	"fmt"
	"github.com/chiquitav2/journalful/internal/api"
	"github.com/chiquitav2/journalful/internal/api/grpc"
	_ "github.com/go-sql-driver/mysql"
	"log/slog"
)

type App struct {
	grpcApi api.ApiModule
	db      *sql.DB
}

func NewApp() *App {
	return &App{}
}
func (s *App) Init() error {
	slog.Info("App initialized")
	// Initialize the database connection
	s.db = connectDB()
	// load the database schema from the embedded schema
	if s.db == nil {
		slog.Error("Database connection is nil, cannot proceed with initialization")
		return fmt.Errorf("failed to connect to the database")
	}

	s.grpcApi = grpc.NewGrpcService(s.db)

	// Start the gRPC server here
	// This is a placeholder for actual server start logic
	if err := s.grpcApi.Register(); err != nil {
		slog.Error("Failed to register gRPC API: %v", err)
		return err
	}
	return nil
}

func (s *App) Start() {
	// Start the server here
	slog.Info("App started")
	// Example: http.ListenAndServe(":8080", nil)
	// This is a placeholder; actual server logic would go here
	if err := s.grpcApi.Start(); err != nil {
		slog.Error("Failed to start gRPC API: %v", err)
	} else {
		slog.Info("gRPC API started successfully")
	}
}

func (s *App) Stop() {
	// Clean up resources here
	if s.db != nil {
		if err := s.db.Close(); err != nil {
			slog.Error("Failed to close database connection: %v", err)
		} else {
			slog.Info("Database connection closed successfully")
		}
	}
	slog.Info("App stopped")
}

func connectDB() *sql.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		"service", "service_password", "localhost", "3306", "journalful")
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		slog.Error("Failed to connect to database: %v", err)
		return nil
	}
	if err = db.Ping(); err != nil {
		slog.Error("Failed to ping database: %v", err)
		return nil
	}
	slog.Info("Connected to database successfully")
	return db
}
