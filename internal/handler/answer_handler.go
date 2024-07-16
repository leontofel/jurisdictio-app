package handler

import (
	"justice-app/internal/model"

	"github.com/gin-gonic/gin"
)

func CreateAnswer(c *gin.Context) {
	var answer model.Answer
	if err := c.ShouldBindJSON(&answer); err != nil {
		log.Error("Failed to bind JSON: ", err)
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	if err := db.Create(&answer).Error; err != nil {
		log.Error("Error creating answer: ", err)
		c.JSON(400, gin.H{"error": "Failed to create answer"})
		return
	}

	c.JSON(200, answer)
}

func GetAnswerByID(c *gin.Context) {
	var answers []model.Answer
	db.Where("author_id = ?", c.Param("id")).Find(&answers)
	c.JSON(200, answers)
}

func GetUserAnswers(c *gin.Context) {
	userID, exists := c.Get("id")
	if !exists {
		c.JSON(401, gin.H{"error": "User ID not found"})
		return
	}

	var answers []model.Answer
	if err := db.Where("author_id = ?", userID).Find(&answers).Error; err != nil {
		log.Error("Error retrieving user answers: ", err)
		c.JSON(500, gin.H{"error": "Failed to retrieve user answers"})
		return
	}

	c.JSON(201, answers)
}