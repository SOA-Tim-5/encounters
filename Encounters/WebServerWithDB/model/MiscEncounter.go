package model

import (
	"github.com/google/uuid"
)

type MiscEncounter struct {
	EncounterId   uuid.UUID `gorm:"primaryKey"`
	Encounter Encounter
	ChallengeDone bool
}