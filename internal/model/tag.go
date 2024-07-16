package model

type Tag struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"type:varchar(50);unique;not null"`
	Questions []Question `json:"questions" gorm:"many2many:question_tags"`
}