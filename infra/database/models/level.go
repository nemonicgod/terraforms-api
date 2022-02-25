package models

import (
	"time"
)

type Level struct {
	ID    uint32 `gorm:"primary_key;auto_increment" json:"-"`
	Level int    `gorm:"size:32;not null;" json:"level"`

	// Associations
	Biomes  []Biome  `gorm:"foreignkey:level_id" json:"biomes"`
	Parcels []Parcel `gorm:"foreignkey:level_id" json:"parcels"`
	Zones   []Zone   `gorm:"foreignkey:level_id" json:"zones"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
}

// TableName ...
func (Level) TableName() string {
	return "level"
}
