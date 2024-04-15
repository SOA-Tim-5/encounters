package repo

import (
	"context"
	"database-example/model"
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

/*
func (repo *EncounterRepository) UpdateEncounter(encounter *model.Encounter) error {

}
*/

func (repo *EncounterRepository) GetEncounter(encounterId int64) (*model.Encounter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	encountersCollection := repo.getEncounterCollection()

	var encounter model.Encounter
	err := encountersCollection.FindOne(ctx, bson.M{"id": encounterId}).Decode(&encounter)
	if err != nil {
		repo.store.logger.Println(err)
		return nil, err
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
	patientsCursor, err := encountersCollection.Find(ctx, bson.M{"status": 0})
	if err != nil {
		repo.store.logger.Println(err)
		return nil, err
	}
	if err = patientsCursor.All(ctx, &encounters); err != nil {
		repo.store.logger.Println(err)
		return nil, err
	}
	return &encounters, nil
}

func (repo *EncounterRepository) FindAll() (*[]model.Encounter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	encountersCollection := repo.getEncounterCollection()

	var encounters []model.Encounter
	patientsCursor, err := encountersCollection.Find(ctx, bson.M{})
	if err != nil {
		repo.store.logger.Println(err)
		return nil, err
	}
	if err = patientsCursor.All(ctx, &encounters); err != nil {
		repo.store.logger.Println(err)
		return nil, err
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

	var encounter model.Encounter
	err := encountersCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&encounter)
	if err != nil {
		repo.store.logger.Println(err)
	}
	return &encounter, nil
}

/*
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
