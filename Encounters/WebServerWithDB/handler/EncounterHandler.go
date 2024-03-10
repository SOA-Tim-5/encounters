package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	var id int64
	vars := mux.Vars(req)
	ids, ok := vars["id"]
	if !ok {
		println("id is missing in parameters")
	}
	id, err = strconv.ParseInt(ids, 10, 64)
	err = handler.EncounterService.ActivateEncounter(id, &touristPosition)
	if err != nil {
		println("Error while activating")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *EncounterHandler) FindTouristProgressByTouristId(writer http.ResponseWriter, req *http.Request) {
	strid := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(strid, 10, 64)
	touristProgress, err := handler.EncounterService.FindTouristProgressByTouristId(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	touristProgressDto := model.TouristProgressDto{
		Xp:    touristProgress.Xp,
		Level: touristProgress.Level,
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(touristProgressDto)
}

func (handler *EncounterHandler) FindAllInRangeOf(writer http.ResponseWriter, req *http.Request) {
	strrange := mux.Vars(req)["range"]
	givenRange, err := strconv.ParseFloat(strrange, 64)
	strLong := mux.Vars(req)["long"]
	userLongitude, err := strconv.ParseFloat(strLong, 64)
	strLat := mux.Vars(req)["lat"]
	userLatitude, err := strconv.ParseFloat(strLat, 64)
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
	id, err := strconv.ParseInt(strid, 10, 64)
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
	id, err := strconv.ParseInt(strid, 10, 64)
	strLong := mux.Vars(req)["long"]
	userLongitude, err := strconv.ParseFloat(strLong, 64)
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

func (handler *EncounterHandler) FindEncounterInstance(writer http.ResponseWriter, req *http.Request) {
	strid := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(strid, 10, 64)
	strencounterid := mux.Vars(req)["encounterId"]
	encounterid, err := strconv.ParseInt(strencounterid, 10, 64)
	instance, err := handler.EncounterService.FindInstanceByUser(id, encounterid)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(instance)
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
