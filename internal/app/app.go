package app

import (
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"
	"time"

	"github.com/chiquitav2/journalful/internal/api"
	"github.com/chiquitav2/journalful/internal/api/grpc"
	"github.com/chiquitav2/journalful/pkg/conf"
	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	grpcApi  api.ApiModule
	db       *sql.DB
	config   *conf.Config
	certFile string
	keyFile  string
}

func NewApp(config *conf.Config, certFile, keyFile string) *App {
	return &App{
		config:   config,
		certFile: certFile,
		keyFile:  keyFile,
	}
}
func (s *App) Init() error {
	slog.Info("initializing application")
	// Initialize the database connection
	db, err := connectDB(s.config.Database)
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}
	s.db = db

	s.grpcApi = grpcapi.NewServer(s.db, s.config, s.certFile, s.keyFile)

	if err := s.grpcApi.Register(); err != nil {
		return fmt.Errorf("failed to register gRPC API: %w", err)
	}
	slog.Info("application initialized successfully")
	return nil
}

func (s *App) Start() error {
	slog.Info("starting application")
	if err := s.grpcApi.Start(s.config); err != nil {
		return fmt.Errorf("failed to start gRPC API: %w", err)
	}
	slog.Info("application started successfully")
	return nil
}

func (s *App) Stop() {
	slog.Info("stopping application")
	// Clean up resources here
	if s.grpcApi != nil {
		s.grpcApi.Stop()
	}
	if s.db != nil {
		if err := s.db.Close(); err != nil {
			slog.Error("failed to close database connection", "error", err)
		} else {
			slog.Info("database connection closed successfully")
		}
	}
	slog.Info("application stopped successfully")
}

func connectDB(dbConfig conf.DatabaseConfig) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}
	// Configure connection pooling
	db.SetMaxOpenConns(25)                 // Max number of open connections
	db.SetMaxIdleConns(5)                  // Max number of idle connections
	db.SetConnMaxLifetime(5 * time.Minute) // Max connection reuse time
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	slog.Info("connected to database successfully")
	return db, nil
}
