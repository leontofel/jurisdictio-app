package routes

import (
	"justice-app/internal/handler"
	"justice-app/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(jwtSecret string) *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello from Gin!"})
	})

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
	r.GET("/home", handler.Home)
	
	authorized := r.Group("/api")
	authorized.Use(middleware.ValidateJWT(jwtSecret))
	{
		// user
		authorized.GET("/users", handler.GetUsers)
		authorized.POST("/users", handler.CreateUser)

		// question
		authorized.POST("/question", handler.CreateQuestion)
		authorized.GET("/question", handler.GetUserQuestions)
		authorized.GET("/question/:id", handler.GetQuestionByID)

		// answer
		authorized.POST("/answer", handler.CreateAnswer)
		authorized.GET("/answer/:id", handler.GetUserAnswers)
		authorized.GET("/answer", handler.GetUserAnswers)
	}

	return r
}
