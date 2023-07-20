package connections

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbCtx context.Context
var dbConn *mongo.Client

func ConnectDatabase() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Failed to Load Environment Variables")
	}
	ctx := context.TODO()

	uri := os.Getenv("MONGODB_ADDRESS")
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverAPIOptions).SetMaxPoolSize(10).SetConnectTimeout(10 * time.Second)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Print(err.Error())
		return
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to connec to connections due to error %s", err.Error())
		return
	}

	dbCtx = ctx
	dbConn = client
}

func CloseDatabase() {
	err := dbConn.Disconnect(dbCtx)
	if err != nil {
		return
	}
}

func GetDatabase() (context.Context, *mongo.Client) {
	return dbCtx, dbConn
}
