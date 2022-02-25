package models

import (
	"time"
)

type Biome struct {
	ID uint32 `gorm:"primary_key;auto_increment" json:"-"`

	Parcels []Parcel `gorm:"foreignkey:id" json:"parcels"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
}

// TableName ...
func (Biome) TableName() string {
	return "biome"
}
