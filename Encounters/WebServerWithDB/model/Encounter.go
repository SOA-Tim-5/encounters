package model

import (
	"gorm.io/gorm"
	"time"
	"github.com/google/uuid"

	"math"

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
}


func (encounter *Encounter) BeforeCreate(scope *gorm.DB) error {
	currentTimestamp := time.Now().UnixNano() / int64(time.Microsecond)
    uniqueID := uuid.New().ID()
    encounter.Id = currentTimestamp + int64(uniqueID)
    return nil
}

func IsInRangeOf(givenrange float64, longitude float64, latitude float64, userLongitude float64, userLatitude float64) bool {
	if(longitude==userLongitude && latitude==userLatitude){
		return true
	}
	var distance=math.Acos(math.Sin(3.14/180*(latitude))*math.Sin(3.14/180*(userLatitude))+math.Cos(3.14/180*(latitude))*math.Cos(3.14/180*userLatitude)*math.Cos(3.14/180*longitude-3.14/180*userLongitude))*6371000


	return distance < givenrange
}
