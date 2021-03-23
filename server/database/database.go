package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBTweets is the mongo tweets collection
var DBTweets *mongo.Collection

// DBUsers is the mongo users collection
var DBUsers *mongo.Collection

// DBLists ist the mongo lists collection
var DBLists *mongo.Collection

// ConnectDB connects to the mongo dat
func ConnectDB(uri string) {
	opt := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), opt)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	db := client.Database("TwitterWall")
	DBTweets = db.Collection("tweets")
	DBUsers = db.Collection("users")
	DBLists = db.Collection("lists")

	// Sets TTL for MongoDB, but value is by default 0. Gets set later
	indexModel := mongo.IndexModel{
		Keys: bson.D{{
			Key:   "expiration",
			Value: 1,
		}},
		Options: options.Index().SetExpireAfterSeconds(0),
	}

	if _, err := DBTweets.Indexes().CreateOne(context.TODO(), indexModel); err != nil {
		log.Fatal(err)
	}
}
