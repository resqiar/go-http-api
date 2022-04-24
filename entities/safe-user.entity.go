package entities

import (
	"time"
)

type SafeUser struct {
	ID        int64
	Username  string
	Bio       string
	Questions []Question `gorm:"foreignKey:AuthorID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
