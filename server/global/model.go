package global

import (
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	ID        uint           `gorm:"primarykey"`
}
