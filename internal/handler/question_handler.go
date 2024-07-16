package handler

import (
	"justice-app/internal/model"

	"github.com/gin-gonic/gin"
)

func CreateQuestion(c *gin.Context) {
	var q model.Question
	if err := c.ShouldBindJSON(&q); err != nil {
		log.Error("Failed to bind JSON: ", err)
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	if err := db.Create(&q).Error; err != nil {
		log.Error("Error creating question: ", err)
		c.JSON(400, gin.H{"error": "Failed to create question"})
		return
	}

	c.JSON(200, q)
}

func GetQuestionByID(c *gin.Context) {
	var questions []model.Question
	db.Where("id = ?", c.Param("id")).Find(&questions)
	c.JSON(200, questions)
}

func GetUserQuestions(c *gin.Context) {
	userID, exists := c.Get("id")
	if !exists {
		c.JSON(401, gin.H{"error": "User ID not found"})
		return
	}

	var questions []model.Question
	if err := db.Where("author_id = ?", userID).Preload("Tags").Find(&questions).Error; err != nil {
		log.Error("Error retrieving user questions: ", err)
		c.JSON(500, gin.H{"error": "Failed to retrieve user questions"})
		return
	}

	c.JSON(201, questions)
}