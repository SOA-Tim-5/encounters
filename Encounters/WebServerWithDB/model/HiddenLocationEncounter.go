package model

import (
	"math"
)

//"github.com/google/uuid"

type HiddenLocationEncounter struct {
	EncounterId      int64 `gorm:"primaryKey"`
	Encounter        Encounter
	PictureLongitude float64
	PictureLatitude  float64
}

type HiddenLocationEncounterDto struct {
	PictureLongitude float64
	PictureLatitude  float64
	Title            string
	Description      string
	Picture          string
	Longitude        float64
	Latitude         float64
	Radius           float64
	XpReward         int
	Status           EncounterStatus
	Type             EncounterType
}

func IsUserInRange(encounter *Encounter, userLongitude float64, userLatitude float64) bool {
	var earthRadius float64 = 6371000
	var latitudeDistance = userLatitude - encounter.Latitude
	var longitudeDistance = userLongitude - encounter.Longitude
	var a = math.Sin(latitudeDistance/2)*math.Sin(latitudeDistance/2) + math.Cos(encounter.Latitude)*math.Cos(userLatitude) + math.Sin(longitudeDistance/2)*math.Sin(longitudeDistance/2)
	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	var distance = earthRadius * c

	return distance < 10
}
