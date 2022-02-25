package models

import (
	"time"
)

type Zone struct {
	ID     uint32 `gorm:"primary_key;auto_increment" json:"-"`
	Name   string `gorm:"size:50;not null;unique" json:"name"`
	Colors string `gorm:"size:255;not null;unique" json:"colors"`

	Parcels []Parcel `gorm:"foreignkey:id" json:"parcels"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
}

// TableName ...
func (Zone) TableName() string {
	return "zone"
}
