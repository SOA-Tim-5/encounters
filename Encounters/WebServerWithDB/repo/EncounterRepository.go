package repo

import (
	"context"
	"database-example/model"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EncounterRepository struct {
	cli    *mongo.Client
	logger *log.Logger
}

func New(ctx context.Context, logger *log.Logger) (*EncounterRepository, error) {
	//dburi := os.Getenv("MONGO_DB_URI")
	dburi := "mongodb://root:pass@mongo:27017"

	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &EncounterRepository{
		cli:    client,
		logger: logger,
	}, nil
}

func (pr *EncounterRepository) Disconnect(ctx context.Context) error {
	err := pr.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (repo *EncounterRepository) getCollection() *mongo.Collection {
	encounterDatabase := repo.cli.Database("mongoDemo")
	encountersCollection := encounterDatabase.Collection("encounters")
	return encountersCollection
}

func (repo *EncounterRepository) CreateMiscEncounter(miscEncounter *model.MiscEncounter) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	encountersCollection := repo.getCollection()

	result, err := encountersCollection.InsertOne(ctx, &miscEncounter)
	if err != nil {
		repo.logger.Println(err)
		return err
	}
	repo.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (repo *EncounterRepository) CreateHiddenLocationEncounter(hiddenLocationEncounter *model.HiddenLocationEncounter) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	encountersCollection := repo.getCollection()

	result, err := encountersCollection.InsertOne(ctx, &hiddenLocationEncounter)
	if err != nil {
		repo.logger.Println(err)
		return err
	}
	repo.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (repo *EncounterRepository) CreateSocialEncounter(socialEncounter *model.SocialEncounter) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	encountersCollection := repo.getCollection()

	result, err := encountersCollection.InsertOne(ctx, &socialEncounter)
	if err != nil {
		repo.logger.Println(err)
		return err
	}
	repo.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

/*
	func (repo *EncounterRepository) UpdateEncounter(encounter *model.Encounter) error {
		dbResult := repo.DatabaseConnection.Save(encounter)
		if dbResult.Error != nil {
			return dbResult.Error
		}
		println("Rows affected: ", dbResult.RowsAffected)
		return nil
	}

func (repo *EncounterRepository) GetEncounter(encounterId int64) (*model.Encounter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	encountersCollection := repo.getCollection()

	var encounter model.Encounter
	err := encountersCollection.FindOne(ctx, bson.M{"_id": encounterId}).Decode(&encounter)
	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	return &encounter, nil
}

func (repo *EncounterRepository) GetHiddenLocationEncounter(encounterId int64) *model.HiddenLocationEncounter {
	var encounter *model.HiddenLocationEncounter
	dbResult := repo.DatabaseConnection.Preload("Encounter").Where("encounter_id = ?", encounterId).First(&encounter)
	if dbResult.Error != nil {
		return nil
	}
	println("Found hidden location encounter")
	return encounter
}

func (repo *EncounterRepository) GetSocialEncounter(encounterId int64) *model.SocialEncounter {
	var encounter *model.SocialEncounter
	dbResult := repo.DatabaseConnection.Preload("Encounter").Where("encounter_id = ?", encounterId).First(&encounter)
	if dbResult.Error != nil {
		return nil
	}
	println("Found social encounter")
	return encounter
}

/*
func (repo *EncounterRepository) FindActiveEncounters() ([]model.Encounter, error) {
	var activeEncounters []model.Encounter
	dbResult := repo.DatabaseConnection.Find(&activeEncounters, "status = 0")
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return activeEncounters, nil
}

func (repo *EncounterRepository) FindAll() ([]model.Encounter, error) {
	var encounters []model.Encounter
	dbResult := repo.DatabaseConnection.Find(&encounters)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return encounters, nil
}

func (repo *EncounterRepository) FindHiddenLocationEncounterById(id int64) (model.HiddenLocationEncounter, error) {
	hiddenLocationEncounter := model.HiddenLocationEncounter{}
	dbResult := repo.DatabaseConnection.First(&hiddenLocationEncounter, "encounter_id = ?", id)
	if dbResult != nil {
		return hiddenLocationEncounter, dbResult.Error
	}
	return hiddenLocationEncounter, nil
}

func (repo *EncounterRepository) FindEncounterById(id int64) (model.Encounter, error) {
	var encounter model.Encounter
	dbResult := repo.DatabaseConnection.First(&encounter, "id=?", id)
	if dbResult != nil {
		return encounter, dbResult.Error
	}

	return encounter, nil
}

func (repo *EncounterRepository) HasUserActivatedOrCompletedEncounter(encounterId int64, userId int64) bool {
	var instance *model.EncounterInstance
	dbResult := repo.DatabaseConnection.Where("encounter_id = ? and user_id = ?", encounterId, userId).First(&instance)
	if dbResult.Error != nil {
		println("Can't be activated")
		return false
	}
	return true
}
*/
