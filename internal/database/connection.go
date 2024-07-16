package database

import (
	"justice-app/internal/config"
	"justice-app/internal/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(conf *config.Config) {
	var err error
	DB, err = gorm.Open(mysql.Open(conf.DBDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
}

func AutoMigrate() {
	DB.AutoMigrate(&model.User{}, &model.Question{}, &model.Tag{}, &model.Answer{})
}