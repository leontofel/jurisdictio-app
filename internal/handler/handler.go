package handler

import (
	"fmt"
	"justice-app/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func Initialize(dbInstance *gorm.DB) {
	db = dbInstance
}

func GetUsers(c *gin.Context) {
	var users []model.User
	err := db.Find(&users).Error
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, users)
	}
}

// create users also take care of errors or nil
func CreateUser(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	err := db.Create(&user).Error
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, user)
	}
}
