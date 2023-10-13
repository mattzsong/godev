package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mattzsong/godev-backend/config"
	"github.com/mattzsong/godev-backend/storage"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	var collections = [2]string{"recipe-lists", "users"}

	env, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Config failed to load: %s\n", err)
		return
	}

	mongodb, err := storage.BootstrapMongo(env.MONGODB_URI, env.MONGODB_NAME, 10*time.Second)
	if err != nil {
		fmt.Printf("Failed to connect to db: %s", err)
		return
	}

	var total int
	for _, coll_name := range collections {
		res, err := mongodb.Collection(coll_name).DeleteMany(context.TODO(), bson.D{})
		if err != nil {
			fmt.Printf("Failed to delete collect %s: %s", coll_name, err)
			return
		}
		total += int(res.DeletedCount)

	}
	fmt.Printf("Successfully removed %d documents from db %s", total, env.MONGODB_NAME)
}
