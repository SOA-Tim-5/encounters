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
		Encounter:     model.Encounter{Title: miscEncounterDto.Title, Description: miscEncounterDto.Description,
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
