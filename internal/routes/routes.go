package routes

import (
	"justice-app/internal/handler"
	"justice-app/internal/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(jwtSecret string) *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		//AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowAllOrigins:  true,
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-CSRF-Token", "Accept", "Content-Length", "X-Requested-With", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Access-Control-Allow-Credentials", "Access-Control-Max-Age", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.OPTIONS("/*cors", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type, X-CSRF-Token")
		c.AbortWithStatus(204)
	})

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
		authorized.GET("/question/unanswered", handler.GetUnansweredQuestions)

		// answer
		authorized.POST("/answer", handler.CreateAnswer)
		authorized.GET("/answer/:id", handler.GetAnswerByID)
		authorized.GET("/answer/question/:id", handler.GetUserAnswersByQuestionID)
		authorized.GET("/answer", handler.GetUserAnswers)
	}

	return r
}
