package model

import (
	"github.com/google/uuid"
)

type KeyPointEncounter struct {
	EncounterId   uuid.UUID `gorm:"primaryKey"`
	Encounter Encounter
	KeyPointId int
}