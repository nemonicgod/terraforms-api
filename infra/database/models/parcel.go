package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Parcel struct {
	ID           uint32          `gorm:"primary_key;auto_increment" json:"-"`
	Name         string          `gorm:"size:50;not null;unique" json:"name"`
	Description  string          `gorm:"size:255;not null;unique" json:"description"`
	AnimationURL string          `gorm:"size:100;not null;unique" json:"animation_url"`
	Image        string          `gorm:"type:text" json:"image"`
	Aspect       decimal.Decimal `gorm:"type:numeric(32,4);not null" json:"aspect"`

	Elevation int `gorm:"not null;" json:"elevation"`

	Seed        int `gorm:"size:32;not null;" json:"seed"`
	XCoordinate int `gorm:"size:32;not null;" json:"x_coordinate"`
	YCoordinate int `gorm:"size:32;not null;" json:"y_coordinate"`

	StuctureSpaceX int `gorm:"not null;" json:"structure_space_x"`
	StuctureSpaceY int `gorm:"not null;" json:"structure_space_y"`
	StuctureSpaceZ int `gorm:"not null;" json:"structure_space_z"`

	// Associations
	Biome   Biome  `gorm:"ForeignKey:biome_id;" json:"biome"`
	BiomeID uint32 `gorm:"size:255;not null;" json:"-"`

	Zone   Zone   `gorm:"ForeignKey:zone_id;" json:"zone"`
	ZoneID uint32 `gorm:"size:255;not null;" json:"-"`

	Level   Level  `gorm:"ForeignKey:zone_id;" json:"level"`
	LevelID uint32 `gorm:"size:255;not null;" json:"-"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
}

func (Parcel) TableName() string {
	return "parcel"
}
