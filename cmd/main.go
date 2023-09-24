package main

import (
	"justice-app/internal/handler"
	"justice-app/internal/middleware"
	"justice-app/internal/model"
	"justice-app/pkg"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Detect if running in Docker
	inDocker := pkg.RunningInDocker()

	// Use different DSN based on the environment
	var dsn string
	if inDocker {
		dsn = os.Getenv("DOCKER_DB_DSN")
	} else {
		dsn = os.Getenv("LOCAL_DB_DSN")
	}

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	db.AutoMigrate(&model.User{})

	// Initialize handler package with db instance
	handler.Initialize(db)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello from Gin!"})
	})

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
	authorized := r.Group("/")
	authorized.Use(middleware.ValidateJWT(os.Getenv("JWT_SECRET")))
	{
		authorized.GET("/users", handler.GetUsers)
		authorized.POST("/users", handler.CreateUser)
		authorized.GET("/home", handler.Home)
	}

	r.Run(":8080")
}
