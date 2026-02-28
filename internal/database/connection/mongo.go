package connection

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"noc-monitoring-bases/internal/config"
	"sync"
)

var (
	clientInstance *mongo.Client
	clientOnce     sync.Once
)

func getMongoClient(uri string) *mongo.Client {
	clientOnce.Do(func() {
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
		if err != nil {
			panic(err)
		}
		clientInstance = client
	})
	return clientInstance
}

func GetCollection(dbName string, collectionName string) *mongo.Collection {
	env, err := config.LoadConfig(".")
	if err != nil {
		_ = fmt.Errorf("error loading config: %v", err)
	}
	return getMongoClient(env.MongoUri).Database(dbName).Collection(collectionName)
}
