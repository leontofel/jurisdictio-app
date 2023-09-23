package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

var db *gorm.DB
var err error

func main() {
	dsn := "user:password@tcp(host.docker.internal:3306)/database_name?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	db.AutoMigrate(&User{})

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello from Gin!"})
	})

	r.GET("/users", GetUsers)
	r.POST("/users", CreateUser)
	r.Run(":8080")
}

func GetUsers(c *gin.Context) {
	var users []User
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
	var user User
	c.BindJSON(&user)
	err := db.Create(&user).Error
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, user)
	}
}
