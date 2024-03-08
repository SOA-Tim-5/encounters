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
	UserId  int64
    Status EncounterInstanceStatus
	CompletionTime time.Time
}

