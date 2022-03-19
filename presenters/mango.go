package presenters

import (
	"backend-score/core/score"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDB struct {
	collection *mongo.Collection
	client     *mongo.Client
}

func NewMangoRepository() score.GameScoreRepository {
	r := mongoDB{}
	InitMango(&r)
	return &r
}

func (db *mongoDB) Disconnect() {
	err := db.client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
}

func (db *mongoDB) Save(s []score.Score) error {
	scoreAsInterface := make([]interface{}, len(s))
	for i, v := range s {
		scoreAsInterface[i] = v
	}
	_, err := db.collection.InsertMany(context.TODO(), scoreAsInterface)
	return err
}

func InitMango(db *mongoDB) {
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

func (db *mongoDB) GetAll() ([]score.Score, error) {

	var scores []score.Score
	cursor, err := db.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return scores, err
	}

	for cursor.Next(context.TODO()) {
		var result score.Score
		if err := cursor.Decode(&result); err != nil {
			return scores, err
		}
		scores = append(scores, result)
	}

	if err := cursor.Err(); err != nil {
		return scores, err
	}

	return scores, nil
}
