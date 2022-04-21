package tasks

import (
	"time"
)

type Task struct {
	ID        int64
	Task      string `gorm:"type:varchar(100);not null"`
	Desc      string
	IsDone    *bool `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
