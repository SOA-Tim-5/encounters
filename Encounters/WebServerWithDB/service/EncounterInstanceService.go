package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type EncounterInstanceService struct {
	EncounterInstanceRepo *repo.EncounterInstanceRepository
}


func (service *EncounterInstanceService) FindInstanceByUser(id int64, encounterid int64) (*model.EncounterInstanceDto, error) {
	instances, err := service.EncounterInstanceRepo.FindInstancesByUserId(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("encounters not found"))
	}
	var foundedInstance model.EncounterInstance
	for _, instance := range instances {

		if instance.EncounterId == encounterid {
			foundedInstance = instance
			break
		}
	}
	instanceDto := model.EncounterInstanceDto{
		UserId:         foundedInstance.UserId,
		Status:         foundedInstance.Status,
		CompletionTime: foundedInstance.CompletionTime,
	}
	return &instanceDto, nil
}
