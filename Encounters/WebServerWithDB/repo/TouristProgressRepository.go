package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type TouristProgressRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TouristProgressRepository) FindTouristProgressByTouristId(id int64) (model.TouristProgress, error) {
	touristProgress := model.TouristProgress{}
	dbResult := repo.DatabaseConnection.First(&touristProgress, "user_id = ?", id)
	if dbResult != nil {
		return touristProgress, dbResult.Error
	}
	return touristProgress, nil
}

func (repo *TouristProgressRepository) UpdateTouristProgress(touristProgress *model.TouristProgress) error {
	dbResult := repo.DatabaseConnection.Save(touristProgress)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *TouristProgressRepository) GetTouristProgress(userId int64) *model.TouristProgress {
	var progress *model.TouristProgress
	dbResult := repo.DatabaseConnection.Where("user_id = ?", userId).First(&progress)
	if dbResult.Error != nil {
		return nil
	}
	return progress
}
