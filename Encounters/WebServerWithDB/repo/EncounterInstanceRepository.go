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


func (repo *EncounterInstanceRepository) CreateEncounterInstance(instance *model.EncounterInstance) error {
	dbResult := repo.DatabaseConnection.Create(instance)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *EncounterInstanceRepository) UpdateEncounterInstance(instance *model.EncounterInstance) error {
	dbResult := repo.DatabaseConnection.Save(instance)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *EncounterInstanceRepository) GetNumberOfActiveInstances(encounterId int64) int64 {
	var instances int64
	dbResult := repo.DatabaseConnection.Where("encounter_id = ? and status = 0", encounterId).Table("encounter_instances").Distinct("user_id").Count(&instances)
	if dbResult.Error != nil {
		return 0
	}
	return instances
}

func (repo *EncounterInstanceRepository) GetActiveInstances(encounterId int64) []*model.EncounterInstance {
	var instances []*model.EncounterInstance
	dbResult := repo.DatabaseConnection.Where("encounter_id = ? and status = 0", encounterId).Find(&instances)
	if dbResult.Error != nil {
		return nil
	}
	return instances
}