package handler

import (
	"justice-app/internal/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateQuestion(c *gin.Context) {
	var q model.Question
	// if err := c.ShouldBindJSON(&q); err != nil {
	// 	log.Error("Failed to bind JSON: ", err)
	// 	c.JSON(400, gin.H{"error": "Invalid request"})
	// 	return
	// }

	// if err := db.Create(&q).Error; err != nil {
	// 	log.Error("Error creating question: ", err)
	// 	c.JSON(400, gin.H{"error": "Failed to create question"})
	// 	return
	// }

	// c.JSON(200, q)

	CreateEntity(c, &q)
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

	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		offset = 0
	}

	var questions []model.Question
	if err := db.Where("author_id = ?", userID).Preload("Tags").Order(c.Param("order")).Offset(offset).Limit(10).Find(&questions).Error; err != nil {
		log.Error("Error retrieving user questions: ", err)
		c.JSON(500, gin.H{"error": "Failed to retrieve user questions"})
		return
	}

	c.JSON(201, questions)
}
func GetUnansweredQuestions(c *gin.Context) {
	offsetStr := c.Query("offset")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	order := c.Query("order")
	if order == "" {
		order = "created_at desc"
	}

	var questions []model.Question
	query := db.Joins("LEFT JOIN answers ON answers.question_id = questions.id").
		Where("answers.id IS NULL").
		Order(order).
		Offset(offset).
		Limit(10)

	if err := query.Find(&questions).Error; err != nil {
		log.Error("Error retrieving unanswered questions: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve unanswered questions"})
		return
	}

	c.JSON(http.StatusOK, questions)
}
