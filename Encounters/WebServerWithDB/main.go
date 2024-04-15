package main

import (
	"context"
	"database-example/handler"
	"database-example/repo"
	"database-example/service"
	"os/signal"

	"log"
	"net/http"
	"os"
	"time"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

/*
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
*/
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

	logger := log.New(os.Stdout, "[encounter-api] ", log.LstdFlags)
	storeLogger := log.New(os.Stdout, "[encounter-store] ", log.LstdFlags)

	store, err := repo.New(timeoutContext, storeLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer store.Disconnect(timeoutContext)

	store.Ping()

	encounterRepo := repo.NewEncounterRepository(store)

	encounterService := service.NewEncounterService(encounterRepo)

	encounterHandler := handler.NewEncounterHandler(encounterService, logger)

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/encounters/misc", encounterHandler.CreateMiscEncounter).Methods("POST")
	router.HandleFunc("/encounters/social", encounterHandler.CreateSocialEncounter).Methods("POST")
	router.HandleFunc("/encounters/hidden", encounterHandler.CreateHiddenLocationEncounter).Methods("POST")

	cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))

	//Initialize the server
	server := http.Server{
		Addr:         ":" + port,
		Handler:      cors(router),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	logger.Println("Server listening on port", port)
	//Distribute all the connections to goroutines
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)

	sig := <-sigCh
	logger.Println("Received terminate, graceful shutdown", sig)

	//Try to shutdown gracefully
	if server.Shutdown(timeoutContext) != nil {
		logger.Fatal("Cannot gracefully shutdown...")
	}
	logger.Println("Server stopped")

}
