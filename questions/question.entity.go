package questions

import (
	"http-api/answers"
	"time"
)

type Question struct {
	ID        int64
	Title     string `gorm:"type:varchar(100);not null"`
	Desc      string
	AuthorID  int64
	Answers   []answers.Answer `gorm:"foreignKey:QuestionID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
