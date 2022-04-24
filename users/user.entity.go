package users

import (
	"http-api/answers"
	"http-api/entities"
	"time"
)

type User struct {
	ID        int64
	Username  string `gorm:"type:varchar(100); unique; not null"`
	Email     string `gorm:"type:varchar(100); unique; not null"`
	Password  string `gorm:"not null"`
	Bio       string
	Questions []entities.Question `gorm:"foreignKey:AuthorID"`
	Answers   []answers.Answer    `gorm:"foreignKey:AuthorID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
