package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type EncounterService struct {
	EncounterRepo         *repo.EncounterRepository
	EncounterInstanceRepo *repo.EncounterInstanceRepository
}

func NewEncounterService(re *repo.EncounterRepository, ri *repo.EncounterInstanceRepository) *EncounterService {
	return &EncounterService{re, ri}
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

/*
	func (service *EncounterService) ActivateEncounter(encounterId int64, position *model.TouristPosition) *model.Encounter {
		var encounter *model.Encounter = service.EncounterRepo.GetEncounter(encounterId)
		fmt.Println("ff %d", position.TouristId)
		if encounter.IsForActivating(position.TouristId, position.Longitude, position.Latitude) && !service.EncounterRepo.HasUserActivatedOrCompletedEncounter(encounterId, position.TouristId) {
			var instance model.EncounterInstance = model.EncounterInstance{
				EncounterId: encounterId, UserId: position.TouristId, Status: model.Activated, CompletionTime: time.Now().UTC(),
			}

			err := service.EncounterInstanceRepo.CreateEncounterInstance(&instance)
			if err != nil {
				return nil
			}
			return encounter
		}
		return nil
	}

	func (service *EncounterService) CompleteHiddenLocationEncounter(encounterId int64, position *model.TouristPosition) error {
		var instance *model.EncounterInstance = service.EncounterInstanceRepo.GetEncounterInstance(encounterId, position.TouristId)
		var encounter *model.HiddenLocationEncounter = service.EncounterRepo.GetHiddenLocationEncounter(encounterId)
		if instance.Status == model.Activated && encounter.IsUserInCompletitionRange(encounter.PictureLongitude, encounter.PictureLatitude, position.Longitude, position.Latitude) {
			instance.Status = model.Completed
			instance.CompletionTime = time.Now().UTC()
			err2 := service.EncounterInstanceRepo.UpdateEncounterInstance(instance)
			if err2 != nil {
				return err2
			}

			var progress *model.TouristProgress = service.TouristProgressRepo.GetTouristProgress(position.TouristId)
			if progress == nil {
				progress = &model.TouristProgress{UserId: position.TouristId, Xp: encounter.Encounter.XpReward, Level: 1}
			} else {
				progress.Xp += encounter.Encounter.XpReward
				progress.Level = progress.Xp/100 + 1
			}
			err3 := service.TouristProgressRepo.UpdateTouristProgress(progress)
			if err3 != nil {
				return err3
			}

		} else {
			println("Can't be completed")
		}
		return nil
	}
*/
func (service *EncounterService) FindAllInRangeOf(givenrange float64, userLongitude float64, userLatitude float64) ([]model.Encounter, error) {
	allencounters, err := service.EncounterRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("encounters not found"))
	}
	var encountersInRange []model.Encounter
	for _, encounter := range *allencounters {
		var d = model.IsInRangeOf(givenrange, encounter.Longitude, encounter.Latitude, userLongitude, userLatitude)
		if d {
			encountersInRange = append(encountersInRange, encounter)
		}
	}

	return encountersInRange, nil
}

func (service *EncounterService) FindAll() (*[]model.Encounter, error) {
	allencounters, err := service.EncounterRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("encounters not found"))
	}

	return allencounters, nil
}

func (service *EncounterService) FindHiddenLocationEncounterById(id int64) (*model.HiddenLocationEncounterDto, error) {
	hiddenLocationEncounter, _ := service.EncounterRepo.FindHiddenLocationEncounterById(id)

	fmt.Printf("%+v\n", hiddenLocationEncounter)
	hiddenLocationEncounterDto := model.HiddenLocationEncounterDto{
		PictureLongitude: hiddenLocationEncounter.PictureLongitude,
		PictureLatitude:  hiddenLocationEncounter.PictureLatitude,
		Id:               hiddenLocationEncounter.Encounter.Id,
		Title:            hiddenLocationEncounter.Encounter.Title,
		Description:      hiddenLocationEncounter.Encounter.Description,
		Picture:          hiddenLocationEncounter.Encounter.Picture,
		Longitude:        hiddenLocationEncounter.Encounter.Longitude,
		Latitude:         hiddenLocationEncounter.PictureLatitude,
		Radius:           hiddenLocationEncounter.Encounter.Radius,
		XpReward:         hiddenLocationEncounter.Encounter.XpReward,
		Status:           hiddenLocationEncounter.Encounter.Status,
		Type:             hiddenLocationEncounter.Encounter.Type,
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

func (service *EncounterService) FindAllDoneByUser(id int64) ([]*model.Encounter, error) {
	instances, err := service.EncounterInstanceRepo.FindInstancesByUserId(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("encounters not found"))
	}
	var encounters []*model.Encounter
	for _, instance := range *instances {
		encounter, _ := service.EncounterRepo.FindEncounterById(instance.EncounterId)
		encounters = append(encounters, encounter)
	}
	return encounters, nil
}

/*
func (service *EncounterService) CompleteMiscEncounter(userid int64, encounterid int64) (*model.TouristProgressDto, error) {
	foundedInstance := service.EncounterInstanceRepo.GetEncounterInstance(encounterid, userid)

	service.EncounterInstanceRepo.UpdateEncounterInstance(model.Complete(foundedInstance))

	touristProgress, err := service.TouristProgressRepo.FindTouristProgressByTouristId(userid)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("tourist progress with userid %s not found", userid))
	}

	encounter, err := service.EncounterRepo.FindEncounterById(encounterid)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("encounter with id %s not found", encounterid))
	}

	var AddedXpTouristProgress = model.AddXp(&touristProgress, encounter.XpReward)

	service.TouristProgressRepo.UpdateTouristProgress(AddedXpTouristProgress)

	touristProgressDto := model.TouristProgressDto{
		Xp:    AddedXpTouristProgress.Xp,
		Level: AddedXpTouristProgress.Level,
	}
	return &touristProgressDto, nil
}

func (service *EncounterService) CompleteSocialEncounter(encounterId int64, position *model.TouristPosition) (*model.TouristProgressDto, error) {
	var instance *model.EncounterInstance = service.EncounterInstanceRepo.GetEncounterInstance(encounterId, position.TouristId)
	var encounter *model.SocialEncounter = service.EncounterRepo.GetSocialEncounter(encounterId)
	var numberOfInstances int64 = service.EncounterInstanceRepo.GetNumberOfActiveInstances(encounterId)
	var progress *model.TouristProgress = service.TouristProgressRepo.GetTouristProgress(position.TouristId)
	if instance.Status == model.Activated && encounter.Encounter.IsInRange(position.Longitude, position.Latitude) && encounter.IsEnoughPeople(int(numberOfInstances)) {
		var err error
		progress, err = service.Complete(instance, progress, position.TouristId, &encounter.Encounter)
		if err != nil {
			return nil, nil
		}
	} else {
		println("Can't be completed")
		return nil, nil
	}
	return &model.TouristProgressDto{Xp: progress.Xp, Level: progress.Level}, nil
}

func (service *EncounterService) Complete(instance *model.EncounterInstance, progress *model.TouristProgress, userId int64, encounter *model.Encounter) (*model.TouristProgress, error) {
	instance.Status = model.Completed
	instance.CompletionTime = time.Now().UTC()
	err2 := service.EncounterInstanceRepo.UpdateEncounterInstance(instance)
	if err2 != nil {
		return nil, err2
	}

	if progress == nil {
		progress = &model.TouristProgress{UserId: userId, Xp: encounter.XpReward, Level: encounter.XpReward/100 + 1}
	} else {
		progress.Xp += encounter.XpReward
		progress.Level = progress.Xp/100 + 1
	}
	err3 := service.TouristProgressRepo.UpdateTouristProgress(progress)
	if err3 != nil {
		return nil, err3
	}

	err4 := service.CompleteAllInRange(encounter.Id)
	if err4 != nil {
		return nil, err4
	}
	return progress, nil
}

func (service *EncounterService) CompleteAllInRange(encounterId int64) error {
	var instances []*model.EncounterInstance = service.EncounterInstanceRepo.GetActiveInstances(encounterId)
	if instances == nil {
		return nil
	}

	for i := 0; i < len(instances); i++ {
		var encounter *model.SocialEncounter = service.EncounterRepo.GetSocialEncounter(encounterId)
		var progress *model.TouristProgress = service.TouristProgressRepo.GetTouristProgress(instances[i].UserId)
		_, err := service.Complete(instances[i], progress, instances[i].UserId, &encounter.Encounter)
		if err != nil {
			return err
		}
	}
	return nil
}
*/
