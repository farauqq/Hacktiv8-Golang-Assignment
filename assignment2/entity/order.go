package entity

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderedAt    time.Time `gorm:"default:now()"`
	CustomerName string    `gorm:"not null"`
	Items        []Item    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
