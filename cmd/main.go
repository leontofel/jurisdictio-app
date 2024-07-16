package main

import (
	"justice-app/internal/config"
	"justice-app/internal/database"
	"justice-app/internal/handler"
	"justice-app/internal/routes"
	"justice-app/pkg"
	"log"
)

func main() {
	pkg.InitLogger()

	// Load configuration
	conf := config.LoadConfig()

	// Connect to database
	database.ConnectDatabase(conf)
	database.AutoMigrate()

	// Initialize handlers
	handler.Initialize(database.DB, pkg.GetLogger())

	// Setup Gin router
	r := routes.SetupRouter(conf.JWTSecret)

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}