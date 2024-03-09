package main

import (
	"database-example/handler"
	"database-example/model"
	"database-example/repo"
	"database-example/service"
	"log"
	"net/http"

	"gorm.io/driver/postgres"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dsn := "user=postgres password=super dbname=explorer-v1 host=localhost port=5432 sslmode=disable search_path=encounters"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		print(err)
		return nil
	}

	err = database.AutoMigrate(&model.Encounter{}, &model.HiddenLocationEncounter{}, &model.SocialEncounter{},
		&model.KeyPointEncounter{}, &model.MiscEncounter{}, &model.TouristProgress{})
	if err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}

	return database
}

func startEncounterServer(handler *handler.EncounterHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/encounters/misc", handler.CreateMiscEncounter).Methods("POST")
	router.HandleFunc("/encounters/social", handler.CreateSocialEncounter).Methods("POST")
	router.HandleFunc("/encounters/hidden", handler.CreateHiddenLocationEncounter).Methods("POST")
	router.HandleFunc("/encounters/touristProgress/{id}", handler.FindTouristProgressByTouristId).Methods("GET")
	router.HandleFunc("/encounters/complete/{id}", handler.CompleteHiddenLocationEncounter).Methods("POST")

	println("Server starting")
	log.Fatal(http.ListenAndServe(":81", router))
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}
	encounterRepo := &repo.EncounterRepository{DatabaseConnection: database}
	encounterService := &service.EncounterService{EncounterRepo: encounterRepo}
	encounterHandler := &handler.EncounterHandler{EncounterService: encounterService}
	startEncounterServer(encounterHandler)
}
