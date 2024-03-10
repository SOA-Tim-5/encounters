package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
	"time"
)

type EncounterService struct {
	EncounterRepo *repo.EncounterRepository
}

func (service *EncounterService) CreateMiscEncounter(miscEncounter *model.MiscEncounter) error {
	miscEncounter.Encounter.Type = model.Misc
	err := service.EncounterRepo.CreateMiscEncounter(miscEncounter)
	if err != nil {
		return err
	}
	return nil
}

func (service *EncounterService) CreateHiddenLocationEncounter(hiddenLocationEncounter *model.HiddenLocationEncounter) error {
	hiddenLocationEncounter.Encounter.Type = model.Hidden
	err := service.EncounterRepo.CreateHiddenLocationEncounter(hiddenLocationEncounter)
	if err != nil {
		return err
	}
	return nil
}

func (service *EncounterService) CreateSocialEncounter(socialEncounter *model.SocialEncounter) error {
	socialEncounter.Encounter.Type = model.Social
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
	var encounter *model.Encounter = service.EncounterRepo.GetEncounter(encounterId)
	if encounter.IsForActivating(position.TouristId, position.Longitude, position.Latitude) && !service.EncounterRepo.HasUserActivatedOrCompletedEncounter(encounterId, position.TouristId) {
		var instance model.EncounterInstance = model.EncounterInstance{
			EncounterId: encounterId, UserId: position.TouristId, Status: model.Activated, CompletionTime: time.Now().UTC(),
		}

		err := service.EncounterRepo.CreateEncounterInstance(&instance)
		if err != nil {
			return err
		}
	}
	return nil
}

func (service *EncounterService) FindTouristProgressByTouristId(id int64) (*model.TouristProgress, error) {
	touristProgress, err := service.EncounterRepo.FindTouristProgressByTouristId(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &touristProgress, nil
}

func (service *EncounterService) FindAllInRangeOf(givenrange float64, userLongitude float64, userLatitude float64) ([]model.Encounter, error) {
	allencounters, err := service.EncounterRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("encounters not found"))
	}
	var encountersInRange []model.Encounter
	for _, encounter := range allencounters {
		var d = model.IsInRangeOf(givenrange, encounter.Longitude, encounter.Latitude, userLongitude, userLatitude)
		if d {
			encountersInRange = append(encountersInRange, encounter)
		}
	}

	return encountersInRange, nil
}

func (service *EncounterService) FindAll() ([]model.Encounter, error) {
	allencounters, err := service.EncounterRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("encounters not found"))
	}

	return allencounters, nil
}

func (service *EncounterService) FindHiddenLocationEncounterById(id int64) (*model.HiddenLocationEncounterDto, error) {
	hiddenLocationEncounter, err := service.EncounterRepo.FindHiddenLocationEncounterById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	hiddenLocationEncounterDto := model.HiddenLocationEncounterDto{
		PictureLongitude: hiddenLocationEncounter.PictureLongitude,
		PictureLatitude:  hiddenLocationEncounter.PictureLatitude,
		Id:               hiddenLocationEncounter.Encounter.Id,
		Title:            hiddenLocationEncounter.Encounter.Title,
		Description:      hiddenLocationEncounter.Encounter.Description,
		Picture:          hiddenLocationEncounter.Encounter.Picture,
		Longitude:        hiddenLocationEncounter.Encounter.Longitude,
		Latitude:         hiddenLocationEncounter.Encounter.Latitude,
		Radius:           hiddenLocationEncounter.Encounter.Radius,
		XpReward:         hiddenLocationEncounter.Encounter.XpReward,
		Status:           hiddenLocationEncounter.Encounter.Status,
		Type:             hiddenLocationEncounter.Encounter.Type,
	}
	return &hiddenLocationEncounterDto, nil
}

func (service *EncounterService) IsUserInCompletitionRange(id int64, userLongitude float64, userLatitude float64) bool {
	hiddenLocationEncounter, _ := service.EncounterRepo.FindHiddenLocationEncounterById(id)
	var isUserInCompletitionRange = model.IsUserInCompletitionRange(hiddenLocationEncounter.PictureLongitude,
		hiddenLocationEncounter.PictureLatitude,
		userLongitude, userLatitude)
	return isUserInCompletitionRange
}

func (service *EncounterService) FindAllDoneByUser(id int64) ([]model.Encounter, error) {
	instances, err := service.EncounterRepo.FindInstancesByUserId(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("encounters not found"))
	}
	var encounters []model.Encounter
	for _, instance := range instances {
		encounter, _ := service.EncounterRepo.FindEncounterById(instance.EncounterId)
		encounters = append(encounters, encounter)
	}
	return encounters, nil
}

func (service *EncounterService) FindInstanceByUser(id int64) (*model.EncounterInstanceDto, error) {
	instance, err := service.EncounterRepo.FindInstanceByUserId(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("instance not found"))
	}
	instanceDto := model.EncounterInstanceDto{
		UserId:         instance.UserId,
		Status:         instance.Status,
		CompletionTime: instance.CompletionTime,
	}
	return &instanceDto, nil
}
