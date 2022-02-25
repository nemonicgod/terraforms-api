package models

import (
	"time"

	"github.com/lib/pq"
)

type Biome struct {
	ID         uint32         `gorm:"primary_key;auto_increment" json:"-"`
	Number     int            `gorm:"size:32;not null;" json:"number"`
	Characters pq.StringArray `gorm:"type:text[]" json:"characters"`

	Parcels []Parcel `gorm:"foreignkey:id" json:"parcels"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
}

// TableName ...
func (Biome) TableName() string {
	return "biome"
}
