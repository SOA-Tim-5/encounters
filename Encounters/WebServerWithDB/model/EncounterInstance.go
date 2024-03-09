package model

import (
	"time"
)
type EncounterInstanceStatus int

const (
	Activated    EncounterInstanceStatus = iota 
	Completed                 
)


type EncounterInstance struct {
	Id      int64
	UserId  int64
    Status EncounterInstanceStatus
	CompletionTime time.Time
}

type EncounterInstanceDto struct {
	UserId  int64
    Status EncounterInstanceStatus
	CompletionTime time.Time
}

func (encounterInstance *EncounterInstance) BeforeCreate(scope *gorm.DB) error {
	currentTimestamp := time.Now().UnixNano() / int64(time.Microsecond)
    uniqueID := uuid.New().ID()
    encounterInstance.Id = currentTimestamp + int64(uniqueID)
    return nil
}