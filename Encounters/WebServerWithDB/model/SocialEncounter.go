package model

import (
	"github.com/google/uuid"
)
type SocialEncounter struct {
	EncounterId   uuid.UUID `gorm:"primaryKey"`
	Encounter Encounter
	PeopleNumber int
}