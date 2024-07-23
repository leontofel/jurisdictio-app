package model

import "time"

type Question struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title     string    `json:"title" gorm:"type:varchar(255);not null"`
	Body      string    `json:"body" gorm:"type:text;not null"`
	AuthorID  uint      `json:"author_id" gorm:"not null"`
	Tags      []Tag     `json:"tags" gorm:"many2many:question_tags"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Question) TableName() string {
	return "questions"
}