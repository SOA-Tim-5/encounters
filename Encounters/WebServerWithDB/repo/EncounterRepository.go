package repo

import (
	"context"
	"database-example/model"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EncounterRepository struct {
	store *Repository
}

func NewEncounterRepository(r *Repository) *EncounterRepository {
	return &EncounterRepository{r}
}

func (repo *EncounterRepository) getEncounterCollection() *mongo.Collection {
	db := repo.store.cli.Database("mongoDemo")
	encountersCollection := db.Collection("encounters")
	return encountersCollection
}
func (repo *EncounterRepository) getEncounterInstanceCollection() *mongo.Collection {
	db := repo.store.cli.Database("mongoDemo")
	encountersCollection := db.Collection("instances")
	return encountersCollection
}

func (repo *EncounterRepository) CreateMiscEncounter(miscEncounter *model.MiscEncounter) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	encountersCollection := repo.getEncounterCollection()
	result, err := encountersCollection.InsertOne(ctx, &miscEncounter)
	if err != nil {
		repo.store.logger.Println(err)
		return err
	}
	repo.store.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (repo *EncounterRepository) CreateHiddenLocationEncounter(hiddenLocationEncounter *model.HiddenLocationEncounter) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	encountersCollection := repo.getEncounterCollection()

	result, err := encountersCollection.InsertOne(ctx, &hiddenLocationEncounter)
	if err != nil {
		repo.store.logger.Println(err)
		return err
	}
	repo.store.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (repo *EncounterRepository) CreateSocialEncounter(socialEncounter *model.SocialEncounter) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	encountersCollection := repo.getEncounterCollection()

	result, err := encountersCollection.InsertOne(ctx, &socialEncounter)
	if err != nil {
		repo.store.logger.Println(err)
		return err
	}
	repo.store.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (repo *EncounterRepository) GetEncounter(encounterId int64) (*model.Encounter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	encountersCollection := repo.getEncounterCollection()

	var encounterResponse model.EncounterResponse
	err := encountersCollection.FindOne(ctx, bson.M{"_id": encounterId}).Decode(&encounterResponse)
	if err != nil {
		repo.store.logger.Println(err)
		return nil, err
	}
	encounter := model.Encounter{
		Id:          encounterResponse.ID,
		Title:       encounterResponse.Encounter.Title,
		Description: encounterResponse.Encounter.Description,
		Picture:     encounterResponse.Encounter.Picture,
		Longitude:   encounterResponse.Encounter.Longitude,
		Latitude:    encounterResponse.Encounter.Latitude,
		Radius:      encounterResponse.Encounter.Radius,
		XpReward:    encounterResponse.Encounter.XpReward,
		Status:      encounterResponse.Encounter.Status,
		Type:        encounterResponse.Encounter.Type,
	}
	return &encounter, nil
}

func (repo *EncounterRepository) GetHiddenLocationEncounter(encounterId int64) (*model.HiddenLocationEncounter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	encountersCollection := repo.getEncounterCollection()

	var encounter model.HiddenLocationEncounter
	err := encountersCollection.FindOne(ctx, bson.M{"_id": encounterId}).Decode(&encounter)
	if err != nil {
		repo.store.logger.Println(err)
		return nil, err
	}
	return &encounter, nil
}

func (repo *EncounterRepository) GetSocialEncounter(encounterId int64) (*model.SocialEncounter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	encountersCollection := repo.getEncounterCollection()

	var encounter model.SocialEncounter
	err := encountersCollection.FindOne(ctx, bson.M{"_id": encounterId}).Decode(&encounter)
	if err != nil {
		repo.store.logger.Println(err)
		return nil, err
	}
	return &encounter, nil
}

func (repo *EncounterRepository) FindActiveEncounters() (*[]model.Encounter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	encountersCollection := repo.getEncounterCollection()

	var encounters []model.Encounter
	cursor, err := encountersCollection.Find(ctx, bson.M{"status": 0})
	if err != nil {
		repo.store.logger.Println(err)
		return nil, err
	}
	for cursor.Next(context.Background()) {
		var encounterResponse model.EncounterResponse
		if err := cursor.Decode(&encounterResponse); err != nil {
			log.Fatal(err)
		}
		encounter := model.Encounter{
			Id:          encounterResponse.ID,
			Title:       encounterResponse.Encounter.Title,
			Description: encounterResponse.Encounter.Description,
			Picture:     encounterResponse.Encounter.Picture,
			Longitude:   encounterResponse.Encounter.Longitude,
			Latitude:    encounterResponse.Encounter.Latitude,
			Radius:      encounterResponse.Encounter.Radius,
			XpReward:    encounterResponse.Encounter.XpReward,
			Status:      encounterResponse.Encounter.Status,
			Type:        encounterResponse.Encounter.Type,
		}
		encounters = append(encounters, encounter)

	}
	return &encounters, nil
}

func (repo *EncounterRepository) FindAll() (*[]model.Encounter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	encountersCollection := repo.getEncounterCollection()

	var encounters []model.Encounter
	cursor, err := encountersCollection.Find(ctx, bson.M{})
	if err != nil {
		repo.store.logger.Println(err)
		return nil, err
	}
	for cursor.Next(context.Background()) {
		var encounterResponse model.EncounterResponse
		if err := cursor.Decode(&encounterResponse); err != nil {
			log.Fatal(err)
		}
		encounter := model.Encounter{
			Id:          encounterResponse.ID,
			Title:       encounterResponse.Encounter.Title,
			Description: encounterResponse.Encounter.Description,
			Picture:     encounterResponse.Encounter.Picture,
			Longitude:   encounterResponse.Encounter.Longitude,
			Latitude:    encounterResponse.Encounter.Latitude,
			Radius:      encounterResponse.Encounter.Radius,
			XpReward:    encounterResponse.Encounter.XpReward,
			Status:      encounterResponse.Encounter.Status,
			Type:        encounterResponse.Encounter.Type,
		}
		encounters = append(encounters, encounter)

	}

	return &encounters, nil
}

func (repo *EncounterRepository) FindHiddenLocationEncounterById(id int64) (*model.HiddenLocationEncounter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	encountersCollection := repo.getEncounterCollection()

	var encounter model.HiddenLocationEncounter
	err := encountersCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&encounter)
	if err != nil {
		repo.store.logger.Println(err)
	}
	return &encounter, nil
}

func (repo *EncounterRepository) FindEncounterById(id int64) (*model.Encounter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	encountersCollection := repo.getEncounterCollection()

	var encounterResponse model.EncounterResponse
	err := encountersCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&encounterResponse)
	if err != nil {
		repo.store.logger.Println(err)
		return nil, err
	}
	encounter := model.Encounter{
		Id:          encounterResponse.ID,
		Title:       encounterResponse.Encounter.Title,
		Description: encounterResponse.Encounter.Description,
		Picture:     encounterResponse.Encounter.Picture,
		Longitude:   encounterResponse.Encounter.Longitude,
		Latitude:    encounterResponse.Encounter.Latitude,
		Radius:      encounterResponse.Encounter.Radius,
		XpReward:    encounterResponse.Encounter.XpReward,
		Status:      encounterResponse.Encounter.Status,
		Type:        encounterResponse.Encounter.Type,
	}
	return &encounter, nil
}

func (repo *EncounterRepository) HasUserActivatedOrCompletedEncounter(encounterId int64, userId int64) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	intsancesCollection := repo.getEncounterInstanceCollection()

	var instance model.EncounterInstance
	err := intsancesCollection.FindOne(ctx, bson.M{"encounterid": encounterId, "userid": userId}).Decode(&instance)
	if err != nil {
		println("Can't be activated")
		return false
	}
	return true
}
