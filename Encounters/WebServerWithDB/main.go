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

	"github.com/golang/protobuf/ptypes"
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
			router.HandleFunc("/encounters/{range}/{long}/{lat}", encounterHandler.FindAllInRangeOf).Methods("GET")
			router.HandleFunc("/encounters", encounterHandler.FindAll).Methods("GET")
			router.HandleFunc("/encounters/complete/{encounterId}/social", encounterHandler.CompleteSocialEncounter).Methods("POST")
			router.HandleFunc("/encounters/complete/{id}", encounterHandler.CompleteHiddenLocationEncounter).Methods("POST")
			router.HandleFunc("/encounters/complete/{userid}/{encounterId}/misc", encounterHandler.CompleteMiscEncounter).Methods("GET")
			router.HandleFunc("/encounters/instance/{id}/{encounterId}/encounter", encounterInstanceHandler.FindEncounterInstance).Methods("GET")
		    router.HandleFunc("/encounters/hidden/{id}", encounterHandler.FindHiddenLocationEncounterById).Methods("GET")
			router.HandleFunc("/encounters/isInRange/{id}/{long}/{lat}", encounterHandler.IsUserInCompletitionRange).Methods("GET")
			router.HandleFunc("/encounters/touristProgress/{id}", touristProgressHandler.FindTouristProgressByTouristId).Methods("GET")
			router.HandleFunc("/encounters/activate/{id}", encounterHandler.ActivateEncounter).Methods("POST")
			router.HandleFunc("/encounters/misc", encounterHandler.CreateMiscEncounter).Methods("POST")
			router.HandleFunc("/encounters/social", encounterHandler.CreateSocialEncounter).Methods("POST")
			router.HandleFunc("/encounters/hidden", encounterHandler.CreateHiddenLocationEncounter).Methods("POST")

			router.HandleFunc("/encounters/{range}/{long}/{lat}", encounterHandler.FindAllInRangeOf).Methods("GET")
			router.HandleFunc("/encounters/doneByUser/{id}", encounterHandler.FindAllDoneByUser).Methods("GET")

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
			Radius: request.Radius, XpReward: int(request.XpReward), Status: model.EncounterStatus(model.Active),
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

func (s Server) CreateSocialEncounter(ctx context.Context, socialEncounterDto *encounter.SocialEncounterCreateDto) (*encounter.EncounterResponseDto, error) {
	println("yyyyyyyyyyyyyyyyyyyyyyyyy")

	id := CreateId()
	newSocialEncounter := model.SocialEncounter{
		EncounterId: id,
		Encounter: model.Encounter{Id: id, Title: socialEncounterDto.Title, Description: socialEncounterDto.Description,
			Picture: socialEncounterDto.Picture, Longitude: socialEncounterDto.Longitude, Latitude: socialEncounterDto.Latitude,
			Radius: socialEncounterDto.Radius, XpReward: int(socialEncounterDto.XpReward), Status: model.EncounterStatus(model.Active),
			Type: 0},
		PeopleNumber: int(socialEncounterDto.PeopleNumber),
	}
	encounterService := service.NewEncounterService(s.EncounterRepo, nil, nil)
	err := encounterService.CreateSocialEncounter(&newSocialEncounter)
	if err != nil {
		println("Error while creating a new social encounter")
		return nil, nil
	}
	return &encounter.EncounterResponseDto{
		Id: id, Title: socialEncounterDto.Title, Description: socialEncounterDto.Description,
		Picture: socialEncounterDto.Picture, Longitude: socialEncounterDto.Longitude, Latitude: socialEncounterDto.Latitude,
		Radius: socialEncounterDto.Radius, XpReward: socialEncounterDto.XpReward, Status: encounter.EncounterResponseDto_EncounterStatus(newSocialEncounter.Encounter.Status),
	}, nil
}

func (s Server) CreateHiddenLocationEncounter(ctx context.Context, hiddenLocationEncounterDto *encounter.HiddenLocationEncounterCreateDto) (*encounter.HiddenLocationEncounterResponseDto, error) {
	println("nnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnn")

	id := CreateId()
	newHiddenLocationEncounter := model.HiddenLocationEncounter{
		EncounterId: id,
		Encounter: model.Encounter{Id: id, Title: hiddenLocationEncounterDto.Title, Description: hiddenLocationEncounterDto.Description,
			Picture: hiddenLocationEncounterDto.Picture, Longitude: hiddenLocationEncounterDto.Longitude, Latitude: hiddenLocationEncounterDto.Latitude,
			Radius: hiddenLocationEncounterDto.Radius, XpReward: int(hiddenLocationEncounterDto.XpReward), Status: model.EncounterStatus(model.Active),
			Type: model.Hidden},
		PictureLongitude: hiddenLocationEncounterDto.PictureLongitude,
		PictureLatitude:  hiddenLocationEncounterDto.PictureLatitude,
	}
	encounterService := service.NewEncounterService(s.EncounterRepo, nil, nil)
	err := encounterService.CreateHiddenLocationEncounter(&newHiddenLocationEncounter)
	if err != nil {
		println("Error while creating a new hidden location encounter")
		return nil, nil
	}
	return &encounter.HiddenLocationEncounterResponseDto{
		Id: id, Title: hiddenLocationEncounterDto.Title, Description: hiddenLocationEncounterDto.Description,
		Picture: hiddenLocationEncounterDto.Picture, Longitude: hiddenLocationEncounterDto.Longitude, Latitude: hiddenLocationEncounterDto.Latitude,
		Radius: hiddenLocationEncounterDto.Radius, XpReward: hiddenLocationEncounterDto.XpReward, PictureLongitude: hiddenLocationEncounterDto.PictureLongitude,
		PictureLatitude: hiddenLocationEncounterDto.PictureLatitude, Status: encounter.HiddenLocationEncounterResponseDto_EncounterStatus(hiddenLocationEncounterDto.Status),
	}, nil
}
func (s Server) FindAllInRangeOf(ctx context.Context, request *encounter.UserPositionWithRange) (*encounter.ListEncounterResponseDto, error) {
	encounterService := service.NewEncounterService(s.EncounterRepo, s.EncounterInstanceRepo, s.TouristProgressRepo)
	encounters, err := encounterService.FindAllInRangeOf(request.Range, request.Longitude, request.Latitude)
	if err != nil {
		return nil, err
	}
	println(len(encounters))
	println(request.Range)
	println(request.Longitude)
	println(request.Latitude)

	var encResponse []*encounter.EncounterResponseDto
	for _, enc := range encounters {
		if enc.Type == model.Misc {
			broj := int64(2)
			encResponse = append(encResponse, &encounter.EncounterResponseDto{
				Id:          enc.Id,
				Title:       enc.Title,
				Description: enc.Description,
				Picture:     enc.Picture,
				Longitude:   enc.Longitude,
				Latitude:    enc.Latitude,
				Radius:      enc.Radius,
				XpReward:    int32(enc.XpReward),
				Status:      encounter.EncounterResponseDto_EncounterStatus(enc.Status),
				Type:        broj,
			})
		} else if enc.Type == model.Hidden {
			broj := int64(1)
			encResponse = append(encResponse, &encounter.EncounterResponseDto{
				Id:          enc.Id,
				Title:       enc.Title,
				Description: enc.Description,
				Picture:     enc.Picture,
				Longitude:   enc.Longitude,
				Latitude:    enc.Latitude,
				Radius:      enc.Radius,
				XpReward:    int32(enc.XpReward),
				Status:      encounter.EncounterResponseDto_EncounterStatus(enc.Status),
				Type:        broj,
			})
		} else {
			broj := int64(0)
			encResponse = append(encResponse, &encounter.EncounterResponseDto{
				Id:          enc.Id,
				Title:       enc.Title,
				Description: enc.Description,
				Picture:     enc.Picture,
				Longitude:   enc.Longitude,
				Latitude:    enc.Latitude,
				Radius:      enc.Radius,
				XpReward:    int32(enc.XpReward),
				Status:      encounter.EncounterResponseDto_EncounterStatus(enc.Status),
				Type:        broj,
			})
		}
	}

	return &encounter.ListEncounterResponseDto{Encounters: encResponse}, nil
}

func (s Server) FindEncounterInstance(ctx context.Context, request *encounter.EncounterInstanceId) (*encounter.EncounterInstanceResponseDto, error) {
	encounterInstanceService := service.NewEncounterInstanceService(s.EncounterInstanceRepo)
	instance, _ := encounterInstanceService.FindInstanceByUser(request.Id, request.EncounterId)
	if instance != nil {
		protoTimestamp, _ := ptypes.TimestampProto(instance.CompletionTime)

		if instance.Status == model.Activated {

			return &encounter.EncounterInstanceResponseDto{
				UserId: instance.UserId, Status: 0,
				CompletitionTime: protoTimestamp}, nil
		} else if instance.Status == model.Completed {
			return &encounter.EncounterInstanceResponseDto{
				UserId: instance.UserId, Status: 1,
				CompletitionTime: protoTimestamp}, nil
		}
	}
	return &encounter.EncounterInstanceResponseDto{
		UserId: 0, Status: 0,
		CompletitionTime: ptypes.TimestampNow()}, nil
}

func (s Server) Activate(ctx context.Context, request *encounter.TouristPosition) (*encounter.EncounterResponseDto, error) {
	touristPosition := model.TouristPosition{
		TouristId: request.TouristId, Longitude: request.Longitude, Latitude: request.Latitude,
	}

	println(touristPosition.TouristId)
	println(touristPosition.Longitude)
	println(touristPosition.Latitude)
	println(request.EncounterId)

	encounterService := service.NewEncounterService(s.EncounterRepo, s.EncounterInstanceRepo, s.TouristProgressRepo)
	enc := encounterService.ActivateEncounter(request.EncounterId, &touristPosition)
	if enc == nil {
		println("Error while activating")
		return nil, nil
	}
	broj := int64(0)

	if enc.Type == model.Misc {
		broj := int64(2)
		return &encounter.EncounterResponseDto{
			Id: enc.Id, Title: enc.Title, Description: enc.Description,
			Picture: enc.Picture, Longitude: enc.Longitude, Latitude: enc.Latitude,
			Radius: enc.Radius, XpReward: int32(enc.XpReward), Status: encounter.EncounterResponseDto_EncounterStatus(enc.Status), Type: broj}, nil
	} else if enc.Type == model.Hidden {
		broj := int64(1)
		return &encounter.EncounterResponseDto{
			Id: enc.Id, Title: enc.Title, Description: enc.Description,
			Picture: enc.Picture, Longitude: enc.Longitude, Latitude: enc.Latitude,
			Radius: enc.Radius, XpReward: int32(enc.XpReward), Status: encounter.EncounterResponseDto_EncounterStatus(enc.Status), Type: broj}, nil
	}

	return &encounter.EncounterResponseDto{
		Id: enc.Id, Title: enc.Title, Description: enc.Description,
		Picture: enc.Picture, Longitude: enc.Longitude, Latitude: enc.Latitude,
		Radius: enc.Radius, XpReward: int32(enc.XpReward), Status: encounter.EncounterResponseDto_EncounterStatus(enc.Status), Type: broj}, nil

}
func (s Server) CompleteMisc(ctx context.Context, request *encounter.EncounterInstanceId) (*encounter.TouristProgress, error) {
	encounterService := service.NewEncounterService(s.EncounterRepo, s.EncounterInstanceRepo, s.TouristProgressRepo)
	touristProgress, _ := encounterService.CompleteMiscEncounter(request.Id, request.EncounterId)
	return &encounter.TouristProgress{
		Xp:    int64(touristProgress.Xp),
		Level: int64(touristProgress.Level)}, nil

}
func (s Server) CompleteSocialEncounter(ctx context.Context, request *encounter.TouristPosition) (*encounter.TouristProgress, error) {
	encounterService := service.NewEncounterService(s.EncounterRepo, s.EncounterInstanceRepo, s.TouristProgressRepo)
	newTouristProgress := model.TouristPosition{
		Longitude: request.Longitude,
		Latitude:  request.Latitude,
		TouristId: request.TouristId}

	touristProgress, _ := encounterService.CompleteSocialEncounter(request.EncounterId, &newTouristProgress)
	return &encounter.TouristProgress{
		Xp:    int64(touristProgress.Xp),
		Level: int64(touristProgress.Level)}, nil

}
func (s Server) CompleteHiddenLocationEncounter(ctx context.Context, request *encounter.TouristPosition) (*encounter.Inrange, error) {
	encounterService := service.NewEncounterService(s.EncounterRepo, s.EncounterInstanceRepo, s.TouristProgressRepo)
	newTouristProgress := model.TouristPosition{
		Longitude: request.Longitude,
		Latitude:  request.Latitude,
		TouristId: request.TouristId}

	_ = encounterService.CompleteHiddenLocationEncounter(request.EncounterId, &newTouristProgress)
	return &encounter.Inrange{In: true}, nil
}
func (s Server) IsUserInCompletitionRange(ctx context.Context, request *encounter.Position) (*encounter.Inrange, error) {
	encounterService := service.NewEncounterService(s.EncounterRepo, s.EncounterInstanceRepo, s.TouristProgressRepo)
	isUserInCompletitionRange := encounterService.IsUserInCompletitionRange(request.Id, request.Longitude, request.Latitude)
	return &encounter.Inrange{
		In: isUserInCompletitionRange}, nil

}
