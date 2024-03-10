package model

//"github.com/google/uuid"

type KeyPointEncounter struct {
	EncounterId int64 `gorm:"primaryKey"`
	Encounter   Encounter
	KeyPointId  int
}
