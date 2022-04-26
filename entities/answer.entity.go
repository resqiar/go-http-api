package entities

import (
	"time"

	"gorm.io/gorm"
)

type Answer struct {
	ID         int64
	AnswerText string `gorm:"type:text;not null"`
	AuthorID   int64
	QuestionID int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Deleted    gorm.DeletedAt
}
