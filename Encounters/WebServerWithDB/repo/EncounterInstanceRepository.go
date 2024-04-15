package repo

import (
	"database-example/model"
	"time"

	"context"

	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EncounterInstanceRepository struct {
	store *Repository
}

func NewEncounterInstanceRepository(r *Repository) *EncounterInstanceRepository {
	return &EncounterInstanceRepository{r}
}

func (repo *EncounterInstanceRepository) getEncounterInstanceCollection() *mongo.Collection {
	db := repo.store.cli.Database("mongoDemo")
	encountersCollection := db.Collection("instances")
	return encountersCollection
}

func (repo *EncounterInstanceRepository) FindInstancesByUserId(id int64) (*[]model.EncounterInstance, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	instancesCollection := repo.getEncounterInstanceCollection()

	var instances []model.EncounterInstance
	cursor, err := instancesCollection.Find(ctx, bson.M{"userid": id})
	if err != nil {
		repo.store.logger.Println(err)
		return nil, err
	}
	for cursor.Next(context.Background()) {
		var encounterInstance model.EncounterInstance
		if err := cursor.Decode(&encounterInstance); err != nil {
			log.Fatal(err)
		}
		instances = append(instances, encounterInstance)

	}

	return &instances, nil
}

func (repo *EncounterInstanceRepository) GetEncounterInstance(encounterId int64, userId int64) (*model.EncounterInstance, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	intsancesCollection := repo.getEncounterInstanceCollection()

	var instance model.EncounterInstance
	err := intsancesCollection.FindOne(ctx, bson.M{"encounterid": encounterId, "userid": userId}).Decode(&instance)
	if err != nil {
		repo.store.logger.Println(err)
		return nil, err
	}
	return &instance, nil
}

/*
func (repo *EncounterInstanceRepository) CreateEncounterInstance(instance *model.EncounterInstance) error {
	dbResult := repo.DatabaseConnection.Create(instance)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *EncounterInstanceRepository) UpdateEncounterInstance(instance *model.EncounterInstance) error {
	dbResult := repo.DatabaseConnection.Save(instance)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *EncounterInstanceRepository) GetNumberOfActiveInstances(encounterId int64) int64 {
	var instances int64
	dbResult := repo.DatabaseConnection.Where("encounter_id = ? and status = 0", encounterId).Table("encounter_instances").Distinct("user_id").Count(&instances)
	if dbResult.Error != nil {
		return 0
	}
	return instances
}

func (repo *EncounterInstanceRepository) GetActiveInstances(encounterId int64) []*model.EncounterInstance {
	var instances []*model.EncounterInstance
	dbResult := repo.DatabaseConnection.Where("encounter_id = ? and status = 0", encounterId).Find(&instances)
	if dbResult.Error != nil {
		return nil
	}
	return instances
}
*/
