package connections

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoService struct {
	Ctx    context.Context
	Client *mongo.Client
}

func ConnectDatabase() (*MongoService, error) {
	ctx := context.TODO()
	uri := os.Getenv("MONGODB_ADDRESS")
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverAPIOptions).SetMaxPoolSize(10).SetConnectTimeout(10 * time.Second)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("database Connection Error: %s", err.Error())
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("database Connection Error: %s", err.Error())
	}

	var MongoConnection = MongoService{
		Ctx:    ctx,
		Client: client,
	}

	return &MongoConnection, nil
}

func (Mongo *MongoService) Close() {
	if Mongo.Client != nil {
		if err := Mongo.Client.Disconnect(Mongo.Ctx); err != nil {
			log.Printf("Database Connection Close Error %s", err.Error())
		} else {
			log.Print("Database Connection Close Success %")
		}
	}
}
