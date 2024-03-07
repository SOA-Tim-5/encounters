package service

import (
	"database-example/model"
	"database-example/repo"
)

type EncounterService struct {
	EncounterRepo *repo.EncounterRepository
}

func (service *EncounterService) CreateMiscEncounter(miscEncounter *model.MiscEncounter) error {
	err := service.EncounterRepo.CreateMiscEncounter(miscEncounter)
	if err != nil {
		return err
	}
	return nil
}

func (service *EncounterService) CreateHiddenLocationEncounter(hiddenLocationEncounter *model.HiddenLocationEncounter) error {
	err := service.EncounterRepo.CreateHiddenLocationEncounter(hiddenLocationEncounter)
	if err != nil {
		return err
	}
	return nil
}

func (service *EncounterService) CreateSocialEncounter(socialEncounter *model.SocialEncounter) error {
	err := service.EncounterRepo.CreateSocialEncounter(socialEncounter)
	if err != nil {
		return err
	}
	return nil
}

func (service *EncounterService) CreateKeyPointEncounter(keyPointEncounter *model.KeyPointEncounter) error {
	err := service.EncounterRepo.CreateKeyPointEncounter(keyPointEncounter)
	if err != nil {
		return err
	}
	return nil
}

func (service *EncounterService) ActivateEncounter(encounterId int64, position *model.TouristPosition) error {
	var encounter model.Encounter = *service.EncounterRepo.GetEncounter(encounterId)
	err := model.Activate(&encounter, position.TouristId, position.Longitude, position.Latitude)
	if err == nil {
		return nil
	}

	err2 := service.EncounterRepo.UpdateEncounter(&encounter)
	if err2 != nil {
		return err2
	}
	return nil
}
