package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type EncounterRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *EncounterRepository) CreateEncounter(encounter *model.Encounter) error {
	dbResult := repo.DatabaseConnection.Create(encounter)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *EncounterRepository) CreateMiscEncounter(miscEncounter *model.MiscEncounter) error {
	dbResult := repo.DatabaseConnection.Create(miscEncounter)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *EncounterRepository) CreateHiddenLocationEncounter(hiddenLocationEncounter *model.HiddenLocationEncounter) error {
	dbResult := repo.DatabaseConnection.Create(hiddenLocationEncounter)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *EncounterRepository) CreateSocialEncounter(socialEncounter *model.SocialEncounter) error {
	dbResult := repo.DatabaseConnection.Create(socialEncounter)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}


func (repo *EncounterRepository) CreateEncounterInstance(instance *model.EncounterInstance) error {
	dbResult := repo.DatabaseConnection.Create(instance)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *EncounterRepository) UpdateEncounter(encounter *model.Encounter) error {
	dbResult := repo.DatabaseConnection.Save(encounter)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *EncounterRepository) UpdateEncounterInstance(instance *model.EncounterInstance) error {
	dbResult := repo.DatabaseConnection.Save(instance)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *EncounterRepository) GetEncounter(encounterId int64) *model.Encounter {
	var encounter *model.Encounter
	dbResult := repo.DatabaseConnection.Where("Id = ?", encounterId).First(&encounter)
	if dbResult.Error != nil {
		return nil
	}
	println("Found encounter")
	return encounter
}

func (repo *EncounterRepository) GetHiddenLocationEncounter(encounterId int64) *model.HiddenLocationEncounter {
	var encounter *model.HiddenLocationEncounter
	dbResult := repo.DatabaseConnection.Preload("Encounter").Where("encounter_id = ?", encounterId).First(&encounter)
	if dbResult.Error != nil {
		return nil
	}
	println("Found hidden location encounter")
	return encounter
}

func (repo *EncounterRepository) GetSocialEncounter(encounterId int64) *model.SocialEncounter {
	var encounter *model.SocialEncounter
	dbResult := repo.DatabaseConnection.Preload("Encounter").Where("encounter_id = ?", encounterId).First(&encounter)
	if dbResult.Error != nil {
		return nil
	}
	println("Found social encounter")
	return encounter
}
/*
func (repo *EncounterRepository) FindTouristProgressByTouristId(id int64) (model.TouristProgress, error) {
	touristProgress := model.TouristProgress{}
	dbResult := repo.DatabaseConnection.First(&touristProgress, "user_id = ?", id)
	if dbResult != nil {
		return touristProgress, dbResult.Error
	}
	return touristProgress, nil
}
*/
func (repo *EncounterRepository) FindActiveEncounters() ([]model.Encounter, error) {
	var activeEncounters []model.Encounter
	dbResult := repo.DatabaseConnection.Find(&activeEncounters, "status = 0")
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return activeEncounters, nil
}

func (repo *EncounterRepository) FindAll() ([]model.Encounter, error) {
	var encounters []model.Encounter
	dbResult := repo.DatabaseConnection.Find(&encounters)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return encounters, nil
}

func (repo *EncounterRepository) FindHiddenLocationEncounterById(id int64) (model.HiddenLocationEncounter, error) {
	hiddenLocationEncounter := model.HiddenLocationEncounter{}
	dbResult := repo.DatabaseConnection.First(&hiddenLocationEncounter, "encounter_id = ?", id)
	if dbResult != nil {
		return hiddenLocationEncounter, dbResult.Error
	}
	return hiddenLocationEncounter, nil
}
/*
func (repo *EncounterRepository) FindInstancesByUserId(id int64) ([]model.EncounterInstance, error) {
	var instances []model.EncounterInstance
	dbResult := repo.DatabaseConnection.Find(&instances, "user_id=?", id)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return instances, nil
}

func (repo *EncounterRepository) FindInstanceByUserId(id int64) (model.EncounterInstance, error) {
	var instance model.EncounterInstance
	dbResult := repo.DatabaseConnection.First(&instance, "user_id=?", id)
	if dbResult != nil {
		return instance, dbResult.Error
	}
	return instance, nil

}
*/
func (repo *EncounterRepository) FindEncounterById(id int64) (model.Encounter, error) {
	var encounter model.Encounter
	dbResult := repo.DatabaseConnection.First(&encounter, "id=?", id)
	if dbResult != nil {
		return encounter, dbResult.Error
	}

	return encounter, nil
}

func (repo *EncounterRepository) HasUserActivatedOrCompletedEncounter(encounterId int64, userId int64) bool {
	var instance *model.EncounterInstance
	dbResult := repo.DatabaseConnection.Where("encounter_id = ? and user_id = ?", encounterId, userId).First(&instance)
	if dbResult.Error != nil {
		println("Can't be activated")
		return false
	}
	return true
}
/*
func (repo *EncounterRepository) GetEncounterInstance(encounterId int64, userId int64) *model.EncounterInstance {
	var instance *model.EncounterInstance
	dbResult := repo.DatabaseConnection.Where("encounter_id = ? and user_id = ?", encounterId, userId).First(&instance)
	if dbResult.Error != nil {
		return nil
	}
	return instance
}

func (repo *EncounterRepository) GetTouristProgress(userId int64) *model.TouristProgress {
	var progress *model.TouristProgress
	dbResult := repo.DatabaseConnection.Where("user_id = ?", userId).First(&progress)
	if dbResult.Error != nil {
		return nil
	}
	return progress
}

func (repo *EncounterRepository) UpdateTouristProgress(touristProgress *model.TouristProgress) error {
	dbResult := repo.DatabaseConnection.Save(touristProgress)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
*/
func (repo *EncounterRepository) GetNumberOfActiveInstances(encounterId int64) int64 {
	var instances int64
	dbResult := repo.DatabaseConnection.Where("encounter_id = ? and status = 0", encounterId).Table("encounter_instances").Distinct("user_id").Count(&instances)
	if dbResult.Error != nil {
		return 0
	}
	return instances
}

func (repo *EncounterRepository) GetActiveInstances(encounterId int64) []*model.EncounterInstance {
	var instances []*model.EncounterInstance
	dbResult := repo.DatabaseConnection.Where("encounter_id = ? and status = 0", encounterId).Find(&instances)
	if dbResult.Error != nil {
		return nil
	}
	return instances
}
