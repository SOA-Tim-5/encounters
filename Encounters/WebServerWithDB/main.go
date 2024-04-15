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
	OSTALO JOS

router.HandleFunc("/encounters/activate/{id}", handler.ActivateEncounter).Methods("POST")
router.HandleFunc("/encounters/complete/{id}", handler.CompleteHiddenLocationEncounter).Methods("POST")
router.HandleFunc("/encounters/complete/{userid}/{encounterId}/misc", handler.CompleteMiscEncounter).Methods("GET")
router.HandleFunc("/encounters/complete/{encounterId}/social", handler.CompleteSocialEncounter).Methods("POST")
*/
func main() {

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
	encounterInstanceRepo := repo.NewEncounterInstanceRepository(store)
	toristProgressRepo := repo.NewTouristProgressRepository(store)

	encounterService := service.NewEncounterService(encounterRepo, encounterInstanceRepo, toristProgressRepo)
	encounterInstanceService := service.NewEncounterInstanceService(encounterInstanceRepo)
	toiristProgressService := service.NewTouristProgressService(toristProgressRepo)

	encounterHandler := handler.NewEncounterHandler(encounterService, logger)
	encounterInstanceHandler := handler.NewEncounterInstanceHandler(encounterInstanceService)
	touristProgressHandler := handler.NewTouristProgressHandler(toiristProgressService)

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/encounters/misc", encounterHandler.CreateMiscEncounter).Methods("POST")
	router.HandleFunc("/encounters/social", encounterHandler.CreateSocialEncounter).Methods("POST")
	router.HandleFunc("/encounters/hidden", encounterHandler.CreateHiddenLocationEncounter).Methods("POST")
	router.HandleFunc("/encounters/isInRange/{id}/{long}/{lat}", encounterHandler.IsUserInCompletitionRange).Methods("GET")
	router.HandleFunc("/encounters/{range}/{long}/{lat}", encounterHandler.FindAllInRangeOf).Methods("GET")
	router.HandleFunc("/encounters", encounterHandler.FindAll).Methods("GET")
	router.HandleFunc("/encounters/hidden/{id}", encounterHandler.FindHiddenLocationEncounterById).Methods("GET")
	router.HandleFunc("/encounters/doneByUser/{id}", encounterHandler.FindAllDoneByUser).Methods("GET")
	router.HandleFunc("/encounters/instance/{id}/{encounterId}/encounter", encounterInstanceHandler.FindEncounterInstance).Methods("GET")
	router.HandleFunc("/encounters/touristProgress/{id}", touristProgressHandler.FindTouristProgressByTouristId).Methods("GET")

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
