package model

import (
	"encoding/json"
	"io"
)

type MiscEncounter struct {
	EncounterId   int64     `bson:"_id,omitempty" json:"EncounterId"`
	Encounter     Encounter `bson:"encounter,omitempty" json:"Encounter"`
	ChallengeDone bool      `bson:"challengedone,omitempty" json:"ChallengeDone"`
}

type MiscEncounterDto struct {
	ChallengeDone bool
	Title         string
	Description   string
	Picture       string
	Longitude     float64
	Latitude      float64
	Radius        float64
	XpReward      int
	Status        EncounterStatus
	Type          EncounterType
}

type MiscEncounters []*MiscEncounter

func (p *MiscEncounters) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *MiscEncounter) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
