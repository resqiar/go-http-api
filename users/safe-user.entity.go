package users

import (
	"http-api/questions"
	"time"
)

type SafeUser struct {
	ID        int64
	Username  string
	Bio       string
	Questions []questions.Question `gorm:"foreignKey:AuthorID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
