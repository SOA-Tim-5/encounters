package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
)

type EncounterHandler struct {
	EncounterService *service.EncounterService
}

func (handler *EncounterHandler) CreateMiscEncounter(writer http.ResponseWriter, req *http.Request) {
	var miscEncounterDto model.MiscEncounterDto
	err := json.NewDecoder(req.Body).Decode(&miscEncounterDto)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	newMiscEncounter := model.MiscEncounter{
		Encounter: model.Encounter{Title: miscEncounterDto.Title, Description: miscEncounterDto.Description,
			Picture: miscEncounterDto.Picture, Longitude: miscEncounterDto.Longitude, Latitude: miscEncounterDto.Latitude,
			Radius: miscEncounterDto.Radius, XpReward: miscEncounterDto.XpReward, Status: miscEncounterDto.Status,
			Type: miscEncounterDto.Type},
		ChallengeDone: miscEncounterDto.ChallengeDone,
	}
	err = handler.EncounterService.CreateMiscEncounter(&newMiscEncounter)
	if err != nil {
		println("Error while creating a new misc encounter")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *EncounterHandler) CreateSocialEncounter(writer http.ResponseWriter, req *http.Request) {
	var socialEncounterDto model.SocialEncounterDto
	err := json.NewDecoder(req.Body).Decode(&socialEncounterDto)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	newSocialEncounter := model.SocialEncounter{
		Encounter: model.Encounter{Title: socialEncounterDto.Title, Description: socialEncounterDto.Description,
			Picture: socialEncounterDto.Picture, Longitude: socialEncounterDto.Longitude, Latitude: socialEncounterDto.Latitude,
			Radius: socialEncounterDto.Radius, XpReward: socialEncounterDto.XpReward, Status: socialEncounterDto.Status,
			Type: socialEncounterDto.Type},
		PeopleNumber: socialEncounterDto.PeopleNumber,
	}
	err = handler.EncounterService.CreateSocialEncounter(&newSocialEncounter)
	if err != nil {
		println("Error while creating a new social encounter")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *EncounterHandler) CreateHiddenLocationEncounter(writer http.ResponseWriter, req *http.Request) {
	var hiddenLocationEncounterDto model.HiddenLocationEncounterDto
	err := json.NewDecoder(req.Body).Decode(&hiddenLocationEncounterDto)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	newHiddenLocationEncounter := model.HiddenLocationEncounter{
		Encounter: model.Encounter{Title: hiddenLocationEncounterDto.Title, Description: hiddenLocationEncounterDto.Description,
			Picture: hiddenLocationEncounterDto.Picture, Longitude: hiddenLocationEncounterDto.Longitude, Latitude: hiddenLocationEncounterDto.Latitude,
			Radius: hiddenLocationEncounterDto.Radius, XpReward: hiddenLocationEncounterDto.XpReward, Status: hiddenLocationEncounterDto.Status,
			Type: hiddenLocationEncounterDto.Type},
		PictureLongitude: hiddenLocationEncounterDto.PictureLongitude,
		PictureLatitude:  hiddenLocationEncounterDto.PictureLatitude,
	}
	err = handler.EncounterService.CreateHiddenLocationEncounter(&newHiddenLocationEncounter)
	if err != nil {
		println("Error while creating a new hidden location encounter")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *EncounterHandler) ActivateEncounter(writer http.ResponseWriter, req *http.Request) {
	var touristPosition model.TouristPosition
	err := json.NewDecoder(req.Body).Decode(&touristPosition)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.EncounterService.ActivateEncounter(1709849423260578, &touristPosition)
	if err != nil {
		println("Error while activating")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
