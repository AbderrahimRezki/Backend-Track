package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to mongodb")

	collection := client.Database("test").Collection("trainers")

	deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	trainers := []interface{}{
		Trainer{"Ash", 10, "Pallet Town"},
		Trainer{"Misty", 10, "Cerulean City"},
		Trainer{"Brock", 15, "Pewter City"},
	}

	// for _, trainer := range trainers {
	// 	insertResult, err := collection.InsertOne(context.TODO(), trainer)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	} else {
	// 		fmt.Println("Inserted document: ", insertResult.InsertedID)
	// 	}
	// }

	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted many documents", insertManyResult.InsertedIDs)

	filter := bson.D{
		{Key: "name", Value: bson.D{
			{Key: "$in", Value: bson.A{"Ash", "Misty"}},
		}},
	}

	update := bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "age", Value: 1},
		}},
	}

	updateResult, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Update document: ", updateResult.MatchedCount, updateResult.ModifiedCount)

	var foundDocument Trainer
	err = collection.FindOne(context.TODO(), filter).Decode(&foundDocument)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found document %+v\n", foundDocument)
	findOptions := options.Find()
	findOptions.SetLimit(5)

	cursor, err := collection.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.TODO()) {
		var doc Trainer
		err := cursor.Decode(&doc)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("document: %+v\n", doc)
	}

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
