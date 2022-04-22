package users

import (
	"http-api/tasks"
	"time"
)

type User struct {
	ID        int64
	Username  string `gorm:"type:varchar(100); unique; not null"`
	Email     string `gorm:"type:varchar(100); unique; not null"`
	Password  string `gorm:"not null"`
	Bio       string
	Tasks     []tasks.Task `gorm:"foreignKey:AuthorID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
