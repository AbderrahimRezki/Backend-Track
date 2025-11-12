package data

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBSingleton struct {
	instance *mongo.Database
}

var dbInstance = &DBSingleton{nil}

type TasksCollectionSingleton struct {
	instance *mongo.Collection
}

type UsersCollectionSingleton struct {
	instance *mongo.Collection
}

var tasksCollection = &TasksCollectionSingleton{nil}
var usersCollection = &UsersCollectionSingleton{nil}

func Setup() (*mongo.Database, error) {
	opts := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}

	db := client.Database("task-manager")
	return db, nil
}

func GetDB() (*mongo.Database, error) {
	if dbInstance.instance != nil {
		return dbInstance.instance, nil
	}

	db, err := Setup()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetTasksCollection() (*mongo.Collection, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}

	if tasksCollection.instance != nil {
		return tasksCollection.instance, nil
	}

	return db.Collection("tasks"), nil
}

func GetUsersCollection() (*mongo.Collection, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}

	if usersCollection.instance != nil {
		return usersCollection.instance, nil
	}

	return db.Collection("users"), nil
}
