package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"

	"passIt/internal/auth"
	"passIt/internal/config"
	"passIt/internal/database"

	// "passIt/internal/store"

	"github.com/redis/go-redis/v9"
)

type Server struct {
	port int

	db     database.Service
	gormDB *gorm.DB
}

func NewServer(ctx context.Context, cfg *config.Config, authClient *auth.Client, redisClient *redis.Client) *http.Server {

	dbService := database.New()
	NewServer := &Server{
		port: cfg.App.Port,

		db:     dbService,
		gormDB: dbService.GetGormDB(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(ctx, authClient, redisClient),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 120 * time.Second,
	}

	return server
}
