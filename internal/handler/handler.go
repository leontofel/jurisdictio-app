package handler

import (
	"justice-app/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var db *gorm.DB
var log *logrus.Logger


func Initialize(database *gorm.DB, logger *logrus.Logger) {
	db = database
	log = logger
}

func CreateEntity(c *gin.Context, entity model.Entity) {
	if err := c.ShouldBindJSON(&entity); err != nil {
		log.Error("Failed to bind JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := db.Create(&entity).Error; err != nil {
		log.Error("Error creating entity: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create entity"})
		return
	}

	c.JSON(http.StatusOK, entity)
}

func GetEntity(c *gin.Context, entity model.Entity, id uint) {
	if err := db.First(&entity, id).Error; err != nil {
		log.Error("Error retrieving entity: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve entity"})
		return
	}

	c.JSON(http.StatusOK, entity)
}