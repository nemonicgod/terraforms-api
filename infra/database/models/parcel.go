package models

import (
	"time"

	"github.com/shopspring/decimal"
)

// Coin ...
type Parcel struct {
	ID           uint32          `gorm:"primary_key;auto_increment" json:"-"`
	Name         string          `gorm:"size:50;not null;unique" json:"name"`
	Description  string          `gorm:"size:255;not null;unique" json:"description"`
	AnimationURL string          `gorm:"size:100;not null;unique" json:"animation_url"`
	Aspect       decimal.Decimal `gorm:"type:numeric(32,4);not null" json:"aspect"`
	Image        string          `gorm:"type:text" json:"image"`

	Elevation int    `gorm:"not null;" json:"elevation"`
	Level     int    `gorm:"not null;" json:"level"`
	Zone      string `gorm:"size:50;not null;unique" json:"zone"`

	XCoordinate int `gorm:"not null;" json:"x_coordinate"`
	YCoordinate int `gorm:"not null;" json:"y_coordinate"`

	StuctureSpaceX int `gorm:"not null;" json:"structure_space_x"`
	StuctureSpaceY int `gorm:"not null;" json:"structure_space_y"`
	StuctureSpaceZ int `gorm:"not null;" json:"structure_space_z"`

	// Relations
	// ZoneColors
	// Characters
	// Attributes

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
}
