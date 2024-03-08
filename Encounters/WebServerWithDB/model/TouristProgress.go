package model

import (
	"gorm.io/gorm"
	"time"
	"github.com/google/uuid"
)

type TouristProgress struct {
	Id      int64
	UserId  int64
	Xp      int
	Level   int

}

func (touristProgress *TouristProgress) BeforeCreate(scope *gorm.DB) error {
	currentTimestamp := time.Now().UnixNano() / int64(time.Microsecond)
    uniqueID := uuid.New().ID()
    touristProgress.Id = currentTimestamp + int64(uniqueID)
    return nil
}


type TouristProgressDto struct {
	Xp      int
	Level   int
}