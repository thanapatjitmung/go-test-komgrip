package repository

import (
	"context"
	"thanapatjitmung/go-test-komgrip/entities"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type logRepositoryImpl struct {
	mongoDb *mongo.Client
}

func NewLogRepositoryImpl(mongoDb *mongo.Client) LogRepository {
	return &logRepositoryImpl{mongoDb: mongoDb}
}

func (r *logRepositoryImpl) Create(logEntities entities.Log) error {
	collection := r.mongoDb.Database("logs_db").Collection("logs_beer")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, logEntities, options.InsertOne())
	if err != nil {
		return err
	}

	return nil
}
