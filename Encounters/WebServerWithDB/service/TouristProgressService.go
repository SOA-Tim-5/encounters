package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type TouristProgressService struct {
	TouristProgressRepo *repo.TouristProgressRepository
}

func NewTouristProgressService(r *repo.TouristProgressRepository) *TouristProgressService {
	return &TouristProgressService{r}
}

func (service *TouristProgressService) FindTouristProgressByTouristId(id int64) (*model.TouristProgress, error) {
	touristProgress, err := service.TouristProgressRepo.FindTouristProgressByTouristId(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	if touristProgress == nil {
		var t model.TouristProgress = model.TouristProgress{Id: CreateId(), UserId: id, Xp: 0, Level: 1}
		err = service.TouristProgressRepo.Create(&t)
		if err != nil {
			return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
		}
		touristProgress, err = service.TouristProgressRepo.FindTouristProgressByTouristId(id)
		if err != nil {
			return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
		}
	}
	return touristProgress, nil
}
