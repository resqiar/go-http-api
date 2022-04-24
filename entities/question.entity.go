package entities

import (
	"time"
)

type Question struct {
	ID        int64
	Title     string `gorm:"type:varchar(100);not null"`
	Desc      string `gorm:"type:text;not null"`
	AuthorID  int64
	Answers   []Answer `gorm:"foreignKey:QuestionID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
