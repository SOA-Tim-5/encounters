package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EncounterStatus int

const (
	Active EncounterStatus = iota
	Draft
	Archived
)

type EncounterType int

const (
	Social EncounterType = iota
	Hidden
	Misc
	KeyPoint
)

type Encounter struct {
	Id          uuid.UUID
	Title       string
	Description string
	Picture     string
	Longitude   float64
	Latitude    float64
	Radius      float64
	XpReward    int
	Status      EncounterStatus
	Type        EncounterType
	Instances   []EncounterInstance `gorm:"type:jsonb;"`
}

func (encounter *Encounter) BeforeCreate(scope *gorm.DB) error {
	encounter.Id = uuid.New()
	return nil
}
