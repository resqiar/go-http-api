package entities

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int64
	Username  string       `gorm:"type:varchar(100);unique;not null"`
	Password  string       `gorm:"not null"`
	Verified  sql.NullBool `gorm:"default:false"`
	Bio       string       `gorm:"type:varchar(100)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
