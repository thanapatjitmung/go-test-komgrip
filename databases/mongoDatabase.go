package databases

import (
	"context"
	"fmt"
	"log"
	"thanapatjitmung/go-test-komgrip/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoDatabase(pctx context.Context, cfg *config.MongoDB) *mongo.Client {
	ctx, cancel := context.WithTimeout(pctx, 20*time.Second)
	defer cancel()
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%d/", cfg.User, cfg.Password, cfg.Host, cfg.Port)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Error:Connect to database error : %v", err)
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("Error: Pinging to database error : %v", err)

	}
	return client
}
