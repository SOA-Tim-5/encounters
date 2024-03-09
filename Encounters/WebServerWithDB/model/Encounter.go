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
	var earthRadius float64 = 6371000
	var latitude1=latitude*3.14/180
	var longitude1=longitude*3.14/180
	var latitude2=userLatitude*3.14/180
	var longitude2=userLongitude*3.14/180
	var latitudeDistance = latitude2 - latitude1
	var longitudeDistance = longitude2 - longitude1
	var a = math.Sin(latitudeDistance/2)*math.Sin(latitudeDistance/2) + math.Cos(latitude1)*math.Cos(latitude2) + math.Sin(longitudeDistance/2)*math.Sin(longitudeDistance/2)
	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	var distance = earthRadius * c

	return distance < givenrange
}
