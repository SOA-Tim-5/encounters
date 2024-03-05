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