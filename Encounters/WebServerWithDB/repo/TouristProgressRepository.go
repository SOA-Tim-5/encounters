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