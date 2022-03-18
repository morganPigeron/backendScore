package repository

import (
	"backend-score/model"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MangoDB struct {
	collection *mongo.Collection
	client     *mongo.Client
}

func NewMangoDB() *MangoDB {
	r := new(MangoDB)
	r.InitMango()
	return r
}

func (db *MangoDB) Disconnect() {
	err := db.client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
}

func (db *MangoDB) InsertMany(a []interface{}) {
	_, err := db.collection.InsertMany(context.TODO(), a)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *MangoDB) InitMango() {
	// set client options
	clientOptions := options.Client().ApplyURI("mongodb://root:example@localhost:27017")

	// connect to mango
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	db.client = client

	// check the connection
	err = db.client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	db.collection = client.Database("game").Collection("score")
}

func (db *MangoDB) GetAll() []model.Score {
	cursor, err := db.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var scores []model.Score
	for cursor.Next(context.TODO()) {
		var result model.Score
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		scores = append(scores, result)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return scores
}
