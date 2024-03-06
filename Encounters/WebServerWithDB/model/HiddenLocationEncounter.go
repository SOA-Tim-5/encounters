package model

import (
	"github.com/google/uuid"
)

type HiddenLocationEncounter struct {
	EncounterId   uuid.UUID `gorm:"primaryKey"`
	Encounter Encounter
	PictureLongitude float64
	PictureLatitude float64
}

type HiddenLocationEncounterDto struct {
	PictureLongitude float64
	PictureLatitude float64
	Title         string
	Description   string
	Picture       string
	Longitude     float64
	Latitude      float64
	Radius        float64
	XpReward      int
	Status        EncounterStatus
	Type          EncounterType
}