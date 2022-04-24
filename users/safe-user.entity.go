package users

import (
	"http-api/entities"
	"time"
)

type SafeUser struct {
	ID        int64
	Username  string
	Bio       string
	Questions []entities.Question `gorm:"foreignKey:AuthorID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
