package handler

import (
	"justice-app/internal/model"

	"github.com/gin-gonic/gin"
)


func GetUsers(c *gin.Context) {
	var users []model.User
	err := db.Find(&users).Error
	if err != nil {
		c.AbortWithError(400, err)
		log.Error("Error getting users: ", err)
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
		c.AbortWithError(400, err)
		log.Error("Error creating user: ", err)
	} else {
		c.JSON(200, user)
	}
}
