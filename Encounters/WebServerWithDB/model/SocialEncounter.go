package model

//"github.com/google/uuid"
type SocialEncounter struct {
	Id           int64     `bson:"_id,omitempty" json:"id"`
	Encounter    Encounter `bson:"encounter,omitempty" json:"encounter"`
	PeopleNumber int       `bson:"peoplenumber,omitempty" json:"peoplenumber"`
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

func (social *SocialEncounter) IsEnoughPeople(numberOfInstances int) bool {
	return numberOfInstances >= social.PeopleNumber
}
