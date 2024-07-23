package model

import "time"

type Answer struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Body       string    `json:"body" gorm:"type:text;not null"`
	QuestionID uint      `json:"question_id" gorm:"not null"`
	AuthorID   uint      `json:"author_id" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Answer) TableName() string {
	return "answers"
}