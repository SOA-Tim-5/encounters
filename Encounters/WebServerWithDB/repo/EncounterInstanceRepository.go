package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type EncounterInstanceRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *EncounterInstanceRepository) FindInstancesByUserId(id int64) ([]model.EncounterInstance, error) {
	var instances []model.EncounterInstance
	dbResult := repo.DatabaseConnection.Find(&instances, "user_id=?", id)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return instances, nil
}