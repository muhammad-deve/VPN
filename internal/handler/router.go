package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammad-deve/vpn/internal/config"
	"github.com/muhammad-deve/vpn/internal/service"
)

func NewRouter(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	// Create service and handler
	svc := service.NewService(cfg)
	handler := NewHandler(svc)

	// Setup routes
	SetupRoutes(router, handler)

	return router
}

func SetupRoutes(router *gin.Engine, handler *Handler) {
	// Health check endpoint
	router.GET("/health", handler.HealthCheck)


	api := router.Group("/api")
	{
		api.GET("/status", handler.HealthCheck)
	}

	vpn := router.Group("/vpn/v2")
	{
		vpn.POST("/start", handler.Start)
		vpn.POST("/stop", handler.Stop)
	}
}
