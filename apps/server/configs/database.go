package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect() error {
	ctx,cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client , err := mongo.Connect(ctx,clientOptions)

	if err != nil {
		return err
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