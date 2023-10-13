package storage

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RecipeStorage struct {
	db *mongo.Database
}

func NewRecipeStorage(db *mongo.Database) *RecipeStorage {
	return &RecipeStorage{
		db: db,
	}
}

func BootstrapMongo(uri, dbName string, timeout time.Duration) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}

	return client.Database(dbName), nil
}

func CloseMongo(db *mongo.Database) error {
	return db.Client().Disconnect(context.Background())
}
