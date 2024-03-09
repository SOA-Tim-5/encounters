package model

type SocialEncounter struct {
	EncounterId  int64 `gorm:"primaryKey"`
	Encounter    Encounter
	PeopleNumber int
}

type SocialEncounterDto struct {
	PeopleNumber int
	Title        string
	Description  string
	Picture      string
	Longitude    float64
	Latitude     float64
	Radius       float64
	XpReward     int
	Status       EncounterStatus
	Type         EncounterType
}
