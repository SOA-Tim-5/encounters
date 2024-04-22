package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TouristProgressHandler struct {
	TouristProgressService *service.TouristProgressService
}

func NewTouristProgressHandler(touristProgressService *service.TouristProgressService) *TouristProgressHandler {
	return &TouristProgressHandler{touristProgressService}
}

func (handler *TouristProgressHandler) FindTouristProgressByTouristId(writer http.ResponseWriter, req *http.Request) {
	println("tourist progress")
	strid := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(strid, 10, 64)
	touristProgress, err := handler.TouristProgressService.FindTouristProgressByTouristId(id)
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
