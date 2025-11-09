package data

import (
	"context"
	"task-manager/models"

	"go.mongodb.org/mongo-driver/bson"
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

var tasksCollection = &TasksCollectionSingleton{nil}

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

func GetAllTasks() ([]*models.Task, error) {
	collection, err := GetTasksCollection()
	if err != nil {
		return nil, err
	}

	var allTasks []*models.Task = make([]*models.Task, 0)
	cur, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var task *models.Task
		cur.Decode(&task)

		allTasks = append(allTasks, task)
	}

	return allTasks, nil
}

func GetTask(id string) (*models.Task, error) {
	collection, err := GetTasksCollection()
	if err != nil {
		return nil, err
	}

	filter := bson.D{{
		Key: "id", Value: id,
	}}

	var task *models.Task
	err = collection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func UpdateTask(id string, updatedTask *models.Task) (*models.Task, error) {
	collection, err := GetTasksCollection()
	if err != nil {
		return nil, err
	}

	filter := bson.D{{
		Key: "id", Value: id,
	}}

	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "title", Value: updatedTask.Title},
				{Key: "description", Value: updatedTask.Description},
			},
		}}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	task, err := GetTask(id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func DeleteTask(id string) (*models.Task, error) {
	collection, err := GetTasksCollection()
	if err != nil {
		return nil, err
	}

	task, err := GetTask(id)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{Key: "id", Value: id}}
	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func PostTask(task *models.Task) (*models.Task, error) {
	collection, err := GetTasksCollection()
	if err != nil {
		return nil, err
	}

	_, err = collection.InsertOne(context.TODO(), task)
	if err != nil {
		return nil, err
	}

	return task, nil
}
