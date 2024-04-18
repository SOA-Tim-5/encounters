package repo

import (
	"context"
	"database-example/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TouristProgressRepository struct {
	store *Repository
}

func (repo *TouristProgressRepository) getTouristProgressCollection() *mongo.Collection {
	db := repo.store.cli.Database("mongoDemo")
	progressCollection := db.Collection("touristProgress")
	return progressCollection
}

func NewTouristProgressRepository(r *Repository) *TouristProgressRepository {
	return &TouristProgressRepository{r}
}
func (repo *TouristProgressRepository) FindTouristProgressByTouristId(id int64) (*model.TouristProgress, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	progressCollection := repo.getTouristProgressCollection()

	var progress model.TouristProgress
	err := progressCollection.FindOne(ctx, bson.M{"userid": id}).Decode(&progress)
	if err != nil {
		repo.store.logger.Println(err)
		return nil, err
	}
	return &progress, nil
}

func (repo *TouristProgressRepository) UpdateTouristProgress(touristProgress *model.TouristProgress) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := repo.getTouristProgressCollection()

	filter := bson.M{"_id": touristProgress.Id}
	update := bson.M{"$set": bson.M{
		"xp":    touristProgress.Xp,
		"level": touristProgress.Level,
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

func (repo *TouristProgressRepository) Create(progress *model.TouristProgress) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	progressCollection := repo.getTouristProgressCollection()
	result, err := progressCollection.InsertOne(ctx, &progress)
	if err != nil {
		repo.store.logger.Println(err)
		return err
	}
	repo.store.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}
