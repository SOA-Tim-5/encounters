package model

//"github.com/google/uuid"

type MiscEncounter struct {
	EncounterId   int64 `gorm:"primaryKey"`
	Encounter     Encounter
	ChallengeDone bool
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
