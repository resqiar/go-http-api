package questions

import (
	"time"
)

type Question struct {
	ID        int64
	Title     string `gorm:"type:varchar(100);not null"`
	Desc      string
	AuthorID  int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
