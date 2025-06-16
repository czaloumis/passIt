package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"

	"passIt/internal/database"
)

type Server struct {
	port int

	db database.Service
	gormDB *gorm.DB
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	dbService := database.New()
	NewServer := &Server{
		port: port,

		db: dbService,
		gormDB: dbService.GetGormDB(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 120 * time.Second,
	}

	return server
}
