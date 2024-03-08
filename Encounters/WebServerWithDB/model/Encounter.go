package model

import (
	"gorm.io/gorm"
	"time"

	"github.com/google/uuid"
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
	Id          int64
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
	currentTimestamp := time.Now().UnixNano() / int64(time.Microsecond)
    uniqueID := uuid.New().ID()
    encounter.Id = currentTimestamp + int64(uniqueID)
    return nil
}
