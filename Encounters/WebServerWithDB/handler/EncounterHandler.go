package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type EncounterHandler struct {
	EncounterService *service.EncounterService
	logger           *log.Logger
}

func NewEncounterHandler(encounterService *service.EncounterService, log *log.Logger) *EncounterHandler {
	return &EncounterHandler{encounterService, log}
}

func CreateId() int64 {
	currentTimestamp := time.Now().UnixNano() / int64(time.Microsecond)
	uniqueID := uuid.New().ID()
	return currentTimestamp + int64(uniqueID)
}

func (handler *EncounterHandler) CreateMiscEncounter(writer http.ResponseWriter, req *http.Request) {
	var miscEncounterDto model.MiscEncounterDto
	err := json.NewDecoder(req.Body).Decode(&miscEncounterDto)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	id := CreateId()
	newMiscEncounter := model.MiscEncounter{
		EncounterId: id,
		Encounter: model.Encounter{Id: id, Title: miscEncounterDto.Title, Description: miscEncounterDto.Description,
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
	id := CreateId()
	newSocialEncounter := model.SocialEncounter{
		EncounterId: id,
		Encounter: model.Encounter{Id: id, Title: socialEncounterDto.Title, Description: socialEncounterDto.Description,
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
	id := CreateId()
	newHiddenLocationEncounter := model.HiddenLocationEncounter{
		EncounterId: id,
		Encounter: model.Encounter{Id: id, Title: hiddenLocationEncounterDto.Title, Description: hiddenLocationEncounterDto.Description,
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
	var id int64
	vars := mux.Vars(req)
	ids, ok := vars["id"]
	if !ok {
		println("id is missing in parameters")
	}
	id, err = strconv.ParseInt(ids, 10, 64)
	encounter := handler.EncounterService.ActivateEncounter(id, &touristPosition)
	if encounter == nil || err != nil {
		println("Error while activating")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(writer).Encode(encounter)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *EncounterHandler) CompleteHiddenLocationEncounter(writer http.ResponseWriter, req *http.Request) {
	var touristPosition model.TouristPosition
	err := json.NewDecoder(req.Body).Decode(&touristPosition)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
	}

	vars := mux.Vars(req)
	id, ok := vars["id"]
	if !ok {
		println("id is missing in parameters")
	}
	encounterId, err := strconv.ParseFloat(id, 64)
	err = handler.EncounterService.CompleteHiddenLocationEncounter(int64(encounterId), &touristPosition)
	if err != nil {
		println("Error while completing hidden location encounter")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *EncounterHandler) FindAllInRangeOf(writer http.ResponseWriter, req *http.Request) {
	println("in range")
	strrange := mux.Vars(req)["range"]
	givenRange, _ := strconv.ParseFloat(strrange, 64)
	strLong := mux.Vars(req)["long"]
	userLongitude, _ := strconv.ParseFloat(strLong, 64)
	strLat := mux.Vars(req)["lat"]
	userLatitude, _ := strconv.ParseFloat(strLat, 64)
	encounters, err := handler.EncounterService.FindAllInRangeOf(givenRange, userLongitude, userLatitude)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(writer).Encode(encounters)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *EncounterHandler) FindAll(writer http.ResponseWriter, req *http.Request) {
	encounters, err := handler.EncounterService.FindAll()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(writer).Encode(encounters)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *EncounterHandler) FindHiddenLocationEncounterById(writer http.ResponseWriter, req *http.Request) {
	strid := mux.Vars(req)["id"]
	id, _ := strconv.ParseInt(strid, 10, 64)
	hiddenLocationEncounter, err := handler.EncounterService.FindHiddenLocationEncounterById(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(hiddenLocationEncounter)
}

func (handler *EncounterHandler) IsUserInCompletitionRange(writer http.ResponseWriter, req *http.Request) {
	strid := mux.Vars(req)["id"]
	id, _ := strconv.ParseInt(strid, 10, 64)
	strLong := mux.Vars(req)["long"]
	userLongitude, _ := strconv.ParseFloat(strLong, 64)
	strLat := mux.Vars(req)["lat"]
	userLatitude, err := strconv.ParseFloat(strLat, 64)
	isUserInCompletitionRange := handler.EncounterService.IsUserInCompletitionRange(id, userLongitude, userLatitude)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(isUserInCompletitionRange)
}

func (handler *EncounterHandler) FindAllDoneByUser(writer http.ResponseWriter, req *http.Request) {
	strid := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(strid, 10, 64)
	encounters, _ := handler.EncounterService.FindAllDoneByUser(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(writer).Encode(encounters)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *EncounterHandler) CompleteMiscEncounter(writer http.ResponseWriter, req *http.Request) {
	struserid := mux.Vars(req)["userid"]
	userid, err := strconv.ParseInt(struserid, 10, 64)
	strencounterid := mux.Vars(req)["encounterId"]
	encounterid, err := strconv.ParseInt(strencounterid, 10, 64)
	touristProgress, err := handler.EncounterService.CompleteMiscEncounter(userid, encounterid)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(touristProgress)
}

func (handler *EncounterHandler) CompleteSocialEncounter(writer http.ResponseWriter, req *http.Request) {
	struserid := mux.Vars(req)["encounterId"]
	encounterId, err := strconv.ParseInt(struserid, 10, 64)

	var touristPosition model.TouristPosition
	err2 := json.NewDecoder(req.Body).Decode(&touristPosition)
	if err2 != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
	}

	touristProgress, err := handler.EncounterService.CompleteSocialEncounter(encounterId, &touristPosition)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(touristProgress)
}
