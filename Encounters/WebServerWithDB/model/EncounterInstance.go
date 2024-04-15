package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type EncounterInstanceStatus int

const (
	Activated EncounterInstanceStatus = iota
	Completed
)

type EncounterInstance struct {
	Id             int64                   `bson:"_id,omitempty" json:"id"`
	EncounterId    int64                   `bson:"encounterid,omitempty" json:"encounterid"`
	UserId         int64                   `bson:"userid,omitempty" json:"userid"`
	Status         EncounterInstanceStatus `bson:"encounterinstancestatus,omitempty" json:"encounterinstancestatus"`
	CompletionTime time.Time               `bson:"completitiontime,omitempty" json:"completitiontime"`
}

func CompleteInstance(instance *EncounterInstance, userId int64) *EncounterInstance {
	instance.Status = Completed
	instance.CompletionTime = time.Now()
	return instance
}

func (r EncounterInstance) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *EncounterInstance) Scan(value interface{}) error {
	if value == nil {
		*r = EncounterInstance{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Scan source is not []byte")
	}
	return json.Unmarshal(bytes, r)
}

type EncounterInstanceDto struct {
	UserId         int64
	Status         EncounterInstanceStatus
	CompletionTime time.Time
}

func Complete(encounterInstance *EncounterInstance) *EncounterInstance {
	encounterInstance.Status = 1
	encounterInstance.CompletionTime = time.Now()
	return encounterInstance
}
