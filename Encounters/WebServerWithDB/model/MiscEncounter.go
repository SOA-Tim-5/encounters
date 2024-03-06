package model

import (
	"github.com/google/uuid"
)

type MiscEncounter struct {
	EncounterId   uuid.UUID `gorm:"primaryKey"`
	Encounter Encounter
	ChallengeDone bool
}

type MiscEncounterDto struct {
	ChallengeDone bool
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
