package handler

import (
	//"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type EncounterInstanceHandler struct {
	EncounterInstanceService *service.EncounterInstanceService
}

func (handler *EncounterInstanceHandler) FindEncounterInstance(writer http.ResponseWriter, req *http.Request) {
	strid := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(strid, 10, 64)
	strencounterid := mux.Vars(req)["encounterId"]
	encounterid, err := strconv.ParseInt(strencounterid, 10, 64)
	instance, err := handler.EncounterInstanceService.FindInstanceByUser(id, encounterid)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(instance)
}