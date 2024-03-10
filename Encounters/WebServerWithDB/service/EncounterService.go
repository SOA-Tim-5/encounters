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

func (service *EncounterService) CompleteHiddenLocationEncounter(encounterId int64, position *model.TouristPosition) error {
	var instance *model.EncounterInstance = service.EncounterRepo.GetEncounterInstance(encounterId, position.TouristId)
	var encounter *model.HiddenLocationEncounter = service.EncounterRepo.GetHiddenLocationEncounter(encounterId)
	if instance.Status == model.Activated && encounter.IsUserInCompletitionRange(encounter.PictureLongitude, encounter.PictureLatitude, position.Longitude, position.Latitude) {
		instance.Status = model.Completed
		instance.CompletionTime = time.Now().UTC()
		err2 := service.EncounterRepo.UpdateEncounterInstance(instance)
		if err2 != nil {
			return err2
		}
	} else {
		println("Can't be completed")
	}
	return nil
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
	encounter, err := service.EncounterRepo.FindEncounterById(hiddenLocationEncounter.EncounterId)

	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	fmt.Printf("%+v\n", hiddenLocationEncounter)
	hiddenLocationEncounterDto := model.HiddenLocationEncounterDto{
		PictureLongitude: hiddenLocationEncounter.PictureLongitude,
		PictureLatitude:  hiddenLocationEncounter.PictureLatitude,
		Id:               encounter.Id,
		Title:            encounter.Title,
		Description:      encounter.Description,
		Picture:          encounter.Picture,
		Longitude:        encounter.Longitude,
		Latitude:         encounter.Latitude,
		Radius:           encounter.Radius,
		XpReward:         encounter.XpReward,
		Status:           encounter.Status,
		Type:             encounter.Type,
	}
	return &hiddenLocationEncounterDto, nil
}

func (service *EncounterService) IsUserInCompletitionRange(id int64, userLongitude float64, userLatitude float64) bool {
	hiddenLocationEncounter, _ := service.EncounterRepo.FindHiddenLocationEncounterById(id)
	var isUserInCompletitionRange = hiddenLocationEncounter.IsUserInCompletitionRange(hiddenLocationEncounter.PictureLongitude,
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

func (service *EncounterService) FindInstanceByUser(id int64, encounterid int64) (*model.EncounterInstanceDto, error) {
	instances, err := service.EncounterRepo.FindInstancesByUserId(id)
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


func (service *EncounterService) CompleteMiscEncounter(userid int64, encounterid int64) (*model.TouristProgressDto, error) {
	instances, err := service.EncounterRepo.FindInstancesByUserId(userid)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("instances not found"))
	}
	var foundedInstance model.EncounterInstance
	for _, instance := range instances {
		if instance.EncounterId == encounterid {
			foundedInstance = instance
			break
		}
	}
	service.EncounterRepo.UpdateEncounterInstance(model.Complete(&foundedInstance))

	touristProgress, err := service.EncounterRepo.FindTouristProgressByTouristId(userid)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("tourist progress with userid %s not found", userid))
	}

	encounter, err := service.EncounterRepo.FindEncounterById(encounterid)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("encounter with id %s not found", encounterid))
	}

	var AddedXpTouristProgress=model.AddXp(&touristProgress,encounter.XpReward)

	service.EncounterRepo.UpdateTouristProgress(AddedXpTouristProgress)

	touristProgressDto := model.TouristProgressDto{
		Xp:    AddedXpTouristProgress.Xp,
		Level: AddedXpTouristProgress.Level,
	}
	return &touristProgressDto, nil
}