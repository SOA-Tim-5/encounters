package main

import (
	"Rest/data"
	"context"
	"database-example/handler"
	"database-example/model"
	"log"
	"net/http"
	"os"
	"time"

	"gorm.io/driver/postgres"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dsn := "user=postgres password=super dbname=explorer host=database port=5432 sslmode=disable search_path=encounters"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		print(err)
		return nil
	}

	err = database.AutoMigrate(&model.Encounter{}, &model.HiddenLocationEncounter{}, &model.SocialEncounter{},
		&model.KeyPointEncounter{}, &model.MiscEncounter{}, &model.TouristProgress{}, &model.EncounterInstance{})
	if err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}

	return database
}

func startEncounterServer(handler *handler.EncounterHandler, touristProgressHandler *handler.TouristProgressHandler,
	encounterInstanceHandler *handler.EncounterInstanceHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/encounters/misc", handler.CreateMiscEncounter).Methods("POST")
	router.HandleFunc("/encounters/social", handler.CreateSocialEncounter).Methods("POST")
	router.HandleFunc("/encounters/hidden", handler.CreateHiddenLocationEncounter).Methods("POST")
	router.HandleFunc("/encounters/activate/{id}", handler.ActivateEncounter).Methods("POST")
	router.HandleFunc("/encounters/touristProgress/{id}", touristProgressHandler.FindTouristProgressByTouristId).Methods("GET")
	router.HandleFunc("/encounters/complete/{id}", handler.CompleteHiddenLocationEncounter).Methods("POST")
	router.HandleFunc("/encounters/{range}/{long}/{lat}", handler.FindAllInRangeOf).Methods("GET")
	router.HandleFunc("/encounters", handler.FindAll).Methods("GET")
	router.HandleFunc("/encounters/hidden/{id}", handler.FindHiddenLocationEncounterById).Methods("GET")
	router.HandleFunc("/encounters/isInRange/{id}/{long}/{lat}", handler.IsUserInCompletitionRange).Methods("GET")
	router.HandleFunc("/encounters/doneByUser/{id}", handler.FindAllDoneByUser).Methods("GET")
	router.HandleFunc("/encounters/instance/{id}/{encounterId}/encounter", encounterInstanceHandler.FindEncounterInstance).Methods("GET")
	router.HandleFunc("/encounters/complete/{userid}/{encounterId}/misc", handler.CompleteMiscEncounter).Methods("GET")
	router.HandleFunc("/encounters/complete/{encounterId}/social", handler.CompleteSocialEncounter).Methods("POST")

	println("Server starting")
	log.Fatal(http.ListenAndServe(":81", router))
}

func main() {
	/*database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}
	encounterRepo := &repo.EncounterRepository{DatabaseConnection: database}
	encounterInstanceRepo := &repo.EncounterInstanceRepository{DatabaseConnection: database}
	touristProgressRepo := &repo.TouristProgressRepository{DatabaseConnection: database}

	encounterService := &service.EncounterService{EncounterRepo: encounterRepo, EncounterInstanceRepo: encounterInstanceRepo,
		TouristProgressRepo: touristProgressRepo}
	encounterInstanceService := &service.EncounterInstanceService{EncounterInstanceRepo: encounterInstanceRepo}
	touristProgressService := &service.TouristProgressService{TouristProgressRepo: touristProgressRepo}

	encounterHandler := &handler.EncounterHandler{EncounterService: encounterService}
	touristProgressHandler := &handler.TouristProgressHandler{TouristProgressService: touristProgressService}
	encounterInstanceHandler := &handler.EncounterInstanceHandler{EncounterInstanceService: encounterInstanceService}
	*/
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "81"
	}

	// Initialize context
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//Initialize the logger we are going to use, with prefix and datetime for every log
	logger := log.New(os.Stdout, "[encounter-api] ", log.LstdFlags)
	storeLogger := log.New(os.Stdout, "[encounter-store] ", log.LstdFlags)

	storeEncounter, err := data.New(timeoutContext, storeLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer storeEncounter.Disconnect(timeoutContext)

	startEncounterServer(encounterHandler, touristProgressHandler, encounterInstanceHandler)
}
