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

func (repo *EncounterInstanceRepository) GetEncounterInstance(encounterId int64, userId int64) *model.EncounterInstance {
	var instance *model.EncounterInstance
	dbResult := repo.DatabaseConnection.Where("encounter_id = ? and user_id = ?", encounterId, userId).First(&instance)
	if dbResult.Error != nil {
		return nil
	}
	return instance
}