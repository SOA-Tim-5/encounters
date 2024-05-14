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

func (repo *EncounterInstanceRepository) CreateEncounterInstance(instance *model.EncounterInstance) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	Collection := repo.getEncounterInstanceCollection()
	result, err := Collection.InsertOne(ctx, &instance)
	if err != nil {
		repo.store.logger.Println(err)
		return err
	}
	repo.store.logger.Printf("Documents ID: %v\n", result.InsertedID)
	filter := bson.M{"_id": instance.Id}
	update := bson.M{"$set": bson.M{
		"completitiontime":        instance.CompletionTime,
		"encounterinstancestatus": 0,
		"userid":                  instance.UserId,
		"encounterid":             instance.EncounterId,
	}}
	t, err := Collection.UpdateOne(ctx, filter, update)
	print(t)
	return nil
}

func (repo *EncounterInstanceRepository) UpdateEncounterInstance(instance *model.EncounterInstance) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := repo.getEncounterInstanceCollection()

	filter := bson.M{"_id": instance.Id}
	update := bson.M{"$set": bson.M{
		"completitiontime":        instance.CompletionTime,
		"encounterinstancestatus": instance.Status,
		"userid":                  instance.UserId,
		"encounterid":             instance.EncounterId,
	}}
	result, err := collection.UpdateOne(ctx, filter, update)
	repo.store.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	repo.store.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		repo.store.logger.Println(err)
		return err
	}
	return nil
}

func (repo *EncounterInstanceRepository) GetNumberOfActiveInstances(encounterId int64) int64 {
	instances, _ := repo.GetActiveInstances(encounterId)
	return int64(len(*instances))

}

func (repo *EncounterInstanceRepository) GetActiveInstances(encounterId int64) (*[]model.EncounterInstance, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	instancesCollection := repo.getEncounterInstanceCollection()

	var instances []model.EncounterInstance
	cursor, err := instancesCollection.Find(ctx, bson.M{"encounterid": encounterId, "encounterinstancestatus": 0})
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
