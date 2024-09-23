package dao

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt *time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"deleted_at" json:"deleted_at"`
}
