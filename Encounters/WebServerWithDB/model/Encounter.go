package model

import (
	"time"

	"gorm.io/gorm"

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

func Complete(enc *Encounter, userId int64, longitude float64, latitude float64) *Encounter {
	var instance *EncounterInstance = nil
	if len(enc.Instances) > 0 {
		for i := 0; i < len(enc.Instances); i++ {
			if enc.Instances[i].Status == Activated && enc.Instances[i].UserId == userId {
				instance = &enc.Instances[i]
				break
			}
		}
		if instance != nil && IsUserInRange(enc, longitude, latitude) {
			CompleteInstance(instance, userId)
			return enc
		}
		println("User is not in 5m range")
	} else {
		println("Encounter not active")
	}
	return nil
}
