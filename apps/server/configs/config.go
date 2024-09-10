package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var Client *mongo.Client

func Init() {
    MongoURI := os.Getenv("MONGO_URI")
    if MongoURI == "" {
        MongoURI = "mongodb://localhost:27017" // default URI
    }


   
}

func ConnectDB() error {
	ctx,cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	client , err := mongo.Connect(ctx,clientOptions)

	if err != nil {
		return err
	} else{
	fmt.Println("Connected to MongoDB!")

	}

	err  = client.Ping(ctx,nil)

	if(err != nil){
		return err
	}

	Client = client

	return nil
}


func Close(){
	if Client != nil {
		ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)
		defer cancel()
		Client.Disconnect(ctx)
	}
}


// ENVS



