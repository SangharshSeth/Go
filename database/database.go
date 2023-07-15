package database

import (
	"context"
	"log"
	"os"

	//"path/filepath"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbCtx context.Context
var dbConn *mongo.Client

func ConnectDatabase() {
	log.Print("In connect")
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Failed to Load Environment Variables")
	}
	//certPath := "/Users/sangharsh/Documents/GitHub/GO/utils/sangharsh_cert.pem/humbalele"
	ctx := context.TODO()
	log.Print("After ctx")
	uri := os.Getenv("MONGODB_ADDRESS")
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverAPIOptions)
	log.Print("Entered before connect")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Print("Failed to connect")
		log.Print(err.Error())
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to connec to database due to error %s", err.Error())
		os.Exit(1)
	}

	dbCtx = ctx
	dbConn = client
}

func GetDatabase() (context.Context, *mongo.Client) {
	return dbCtx, dbConn
}
