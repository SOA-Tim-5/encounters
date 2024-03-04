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
	UserId  int
    Status EncounterInstanceStatus
	CompletionTime time.Time
}

