package users

import (
	"http-api/tasks"
	"time"
)

type SafeUser struct {
	ID        int64
	Username  string
	Bio       string
	Tasks     []tasks.Task `gorm:"foreignKey:AuthorID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
