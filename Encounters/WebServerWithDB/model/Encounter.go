package model

import (
	"math"
	"slices"
	"time"

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

func Activate(encounter *Encounter, userId int64, userLongitude float64, userLatitude float64) *Encounter {
	if encounter.Status != Active {
		return nil
	}
	if HasUserActivatedEncounter(encounter, userId) {
		return nil
	}
	if IsInRange(encounter.Radius, encounter.Longitude, encounter.Latitude, userLongitude, userLatitude) {
		return nil
	}

	encounter.Status = Active
	instance := EncounterInstance{
		UserId: int(userId), Status: Activated, CompletionTime: time.Now(),
	}
	encounter.Instances = append(encounter.Instances, instance)
	return encounter
}

func HasUserActivatedEncounter(encounter *Encounter, userId int64) bool {
	return slices.IndexFunc(encounter.Instances, func(c EncounterInstance) bool { return c.Status == Activated }) == 1
}

func IsInRange(radius float64, longitude float64, latitude float64, userLongitude float64, userLatitude float64) bool {
	var earthRadius float64 = 6371000
	var latitudeDistance = userLatitude - latitude
	var longitudeDistance = userLongitude - longitude
	var a = math.Sin(latitudeDistance/2)*math.Sin(latitudeDistance/2) + math.Cos(latitude)*math.Cos(userLatitude) + math.Sin(longitudeDistance/2)*math.Sin(longitudeDistance/2)
	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	var distance = earthRadius * c

	return distance < radius
}
