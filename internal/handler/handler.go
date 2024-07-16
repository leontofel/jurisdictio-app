package handler

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var db *gorm.DB
var log *logrus.Logger


func Initialize(database *gorm.DB, logger *logrus.Logger) {
	db = database
	log = logger
}

