package main

import (
	"github.com/muhammad-deve/vpn/internal/config"
	"github.com/muhammad-deve/vpn/internal/handler"
)

func main() {
	cfg := config.NewConfig()

	// Create and run the Gin router
	router := handler.NewRouter(cfg)

	// Run the server on port 8080
	router.Run(":8080")
}
