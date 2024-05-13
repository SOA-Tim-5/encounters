package main

import (
	"context"
	"database-example/model"
	"database-example/proto/encounter"
	"database-example/repo"
	"database-example/service"
	"log"
	"net"
	"os"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	/*
		port := os.Getenv("PORT")
		if len(port) == 0 {
			port = "81"
		}
	*/
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

	//encounterService := service.NewEncounterService(encounterRepo, nil, nil)
	//encounterInstanceService := service.NewEncounterInstanceService(encounterInstanceRepo)
	//toiristProgressService := service.NewTouristProgressService(toristProgressRepo)

	/*
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
			router.HandleFunc("/encounters/complete/{userid}/{encounterId}/misc", encounterHandler.CompleteMiscEncounter).Methods("GET")
			router.HandleFunc("/encounters/activate/{id}", encounterHandler.ActivateEncounter).Methods("POST")
			router.HandleFunc("/encounters/complete/{id}", encounterHandler.CompleteHiddenLocationEncounter).Methods("POST")
			router.HandleFunc("/encounters/complete/{encounterId}/social", encounterHandler.CompleteSocialEncounter).Methods("POST")

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

	*/

	lis, err := net.Listen("tcp", "localhost:81")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	encounter.RegisterEncounterServer(grpcServer, Server{EncounterRepo: encounterRepo, EncounterInstanceRepo: encounterInstanceRepo, TouristProgressRepo: toristProgressRepo})
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)

}

type Server struct {
	encounter.UnimplementedEncounterServer
	EncounterRepo         *repo.EncounterRepository
	EncounterInstanceRepo *repo.EncounterInstanceRepository
	TouristProgressRepo   *repo.TouristProgressRepository
}

func CreateId() int64 {
	currentTimestamp := time.Now().UnixNano() / int64(time.Microsecond)
	uniqueID := uuid.New().ID()
	return currentTimestamp + int64(uniqueID)
}

func (s Server) CreateMiscEncounter(ctx context.Context, request *encounter.MiscEncounterCreateDto) (*encounter.MiscEncounterResponseDto, error) {
	println("dhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh")
	id := CreateId()
	newMiscEncounter := model.MiscEncounter{
		EncounterId: id,
		Encounter: model.Encounter{Id: id, Title: request.Title, Description: request.Description,
			Picture: request.Picture, Longitude: request.Longitude, Latitude: request.Latitude,
			Radius: request.Radius, XpReward: int(request.XpReward), Status: model.EncounterStatus(request.Status),
			Type: model.Misc},
		ChallengeDone: request.ChallengeDone,
	}
	encounterService := service.NewEncounterService(s.EncounterRepo, nil, nil)
	err := encounterService.CreateMiscEncounter(&newMiscEncounter)
	if err != nil {
		println("Error while creating a new misc encounter")
	}

	return &encounter.MiscEncounterResponseDto{
		Id: id, Title: request.Title, Description: request.Description,
		Picture: request.Picture, Longitude: request.Longitude, Latitude: request.Latitude,
		Radius: request.Radius, XpReward: request.XpReward, Status: encounter.MiscEncounterResponseDto_EncounterStatus(request.Status),
		ChallengeDone: request.ChallengeDone,
	}, nil
}
