package model

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
)

type HiddenLocationEncounter struct {
	EncounterId      int64     `bson:"_id,omitempty" json:"EncounterId"`
	Encounter        Encounter `bson:"encounter,omitempty" json:"Encounter"`
	PictureLongitude float64   `bson:"picturelongitude,omitempty" json:"PictureLongitude"`
	PictureLatitude  float64   `bson:"picturelatitude,omitempty" json:"PictureLatitude"`
}

type HiddenLocationEncounterDto struct {
	PictureLongitude float64
	PictureLatitude  float64
	Id               int64
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
	var distance = math.Acos(math.Sin(3.14/180*(encounter.Latitude))*math.Sin(3.14/180*(userLatitude))+math.Cos(3.14/180*(encounter.Latitude))*math.Cos(3.14/180*userLatitude)*math.Cos(3.14/180*encounter.Longitude-3.14/180*userLongitude)) * 6371000

	return distance < encounter.Radius
}

func (encounter *HiddenLocationEncounter) IsUserInCompletitionRange(longitude float64, latitude float64, userLongitude float64, userLatitude float64) bool {
	if longitude == userLongitude && latitude == userLatitude {
		return true
	}
	var distance = math.Acos(math.Sin(3.14/180*(latitude))*math.Sin(3.14/180*(userLatitude))+math.Cos(3.14/180*(latitude))*math.Cos(3.14/180*userLatitude)*math.Cos(3.14/180*longitude-3.14/180*userLongitude)) * 6371000
	fmt.Println(latitude, longitude, userLatitude, userLongitude)

	fmt.Println(distance)

	return distance < 100
}

type HiddenLocationEncounters []*HiddenLocationEncounter

func (p *HiddenLocationEncounters) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *HiddenLocationEncounter) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
