package database

import (
	"context"
	"log"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database
var dbCtx = context.TODO()

func Init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(dbCtx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(dbCtx, nil)
	if err != nil {
		log.Fatal(err)
	}

	color.Green("⛁ Connected to Database")
	db = client.Database("company_tasks")
}
