package service

import (
	"database-example/model"
	"database-example/repo"
)

type EncounterService struct {
	EncounterRepo *repo.EncounterRepository
}

func (service *EncounterService) CreateEncounter(encounter *model.Encounter) error {
	err := service.EncounterRepo.CreateEncounter(encounter)
	if err != nil {
		return err
	}
	return nil
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