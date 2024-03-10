package model

type EncounterActivationDto struct {
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
