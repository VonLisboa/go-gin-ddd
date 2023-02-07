package database

import (
	"context"
	"log"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mongo *mongo.Database
var MongoContext = context.TODO()

func Init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(MongoContext, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(MongoContext, nil)
	if err != nil {
		log.Fatal(err)
	}

	color.Green("⛁ Connected to Database")
	Mongo = client.Database("company_tasks")
}
