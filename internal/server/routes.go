package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	r.Use(s.rateLimit())

	r.GET("/", s.HelloWorldHandler)

	r.GET("/health", s.healthHandler)

	r.GET("/vehicles/decode/:vin", s.VehicleHandler.DecodeVehicle)
	r.GET("/vehicles/:vin", s.VehicleHandler.CreateVehicle)
	r.POST("/vehicles", s.VehicleHandler.GetVehicle)

	r.GET("/orgs", s.OrgHandler.GetAllOrganisations)
	r.GET("/orgs/:id", s.OrgHandler.GetOrganisation)
	r.POST("/orgs", s.OrgHandler.CreateOrgansation)
	r.PATCH("/orgs", s.OrgHandler.UpdateOrganisation)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	err := s.db.Ping(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf(fmt.Sprintf("db down: %v", err)) // Log the error and terminate the program
		c.JSON(http.StatusInternalServerError, stats)
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"
	c.JSON(http.StatusOK, stats)
}
