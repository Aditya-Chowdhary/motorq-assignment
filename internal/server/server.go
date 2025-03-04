package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"

	"motorq-assignment/internal/controllers/organisations"
	"motorq-assignment/internal/controllers/vehicles"
	"motorq-assignment/internal/database"
)

type config struct {
	limiter struct{
		rate int
		burst int
	}
}

type Server struct {
	port int

	db             *pgxpool.Pool
	OrgHandler     *organisations.OrgHandler
	VehicleHandler *vehicles.VehicleHandler

	config
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	db := database.NewService()
	NewServer := &Server{
		port: port,
		db:   db,

		OrgHandler:     organisations.Handler(db),
		VehicleHandler: vehicles.Handler(db),

		config: config{
			limiter: struct{rate int; burst int}{1, 2},
		},
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf("localhost:%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
