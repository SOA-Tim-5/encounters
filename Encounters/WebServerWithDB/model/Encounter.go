package model

import (
	"encoding/json"
	"io"
	"math"
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
	Id          int64           `bson:"_id,omitempty" json:"id"`
	Title       string          `bson:"title,omitempty" json:"title"`
	Description string          `bson:"description,omitempty" json:"description"`
	Picture     string          `bson:"picture,omitempty" json:"picture"`
	Longitude   float64         `bson:"longitude,omitempty" json:"longitude"`
	Latitude    float64         `bson:"latitude,omitempty" json:"latitude"`
	Radius      float64         `bson:"radius,omitempty" json:"radius"`
	XpReward    int             `bson:"xpreward,omitempty" json:"xpreawrd"`
	Status      EncounterStatus `bson:"status,omitempty" json:"status"`
	Type        EncounterType   `bson:"type,omitempty" json:"type"`
}

func (encounter *Encounter) BeforeCreate(scope *gorm.DB) error {
	currentTimestamp := time.Now().UnixNano() / int64(time.Microsecond)
	uniqueID := uuid.New().ID()
	encounter.Id = currentTimestamp + int64(uniqueID)
	return nil
}

func IsInRangeOf(givenrange float64, longitude float64, latitude float64, userLongitude float64, userLatitude float64) bool {
	if longitude == userLongitude && latitude == userLatitude {
		return true
	}
	var distance = math.Acos(math.Sin(3.14/180*(latitude))*math.Sin(3.14/180*(userLatitude))+math.Cos(3.14/180*(latitude))*math.Cos(3.14/180*userLatitude)*math.Cos(3.14/180*longitude-3.14/180*userLongitude)) * 6371000

	return distance < givenrange
}

func (encounter *Encounter) IsForActivating(userId int64, userLongitude float64, userLatitude float64) bool {
	return encounter.Status == Active && encounter.IsInRange(userLongitude, userLatitude)
}

func (encounter *Encounter) IsInRange(userLongitude float64, userLatitude float64) bool {
	var distance = math.Acos(math.Sin(3.14/180*(encounter.Latitude))*math.Sin(3.14/180*(userLatitude))+math.Cos(3.14/180*(encounter.Latitude))*math.Cos(3.14/180*userLatitude)*math.Cos(3.14/180*encounter.Longitude-3.14/180*userLongitude)) * 6371000
	return distance < encounter.Radius
}

type Encounters []*Encounter

func (p *Encounters) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Encounter) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
