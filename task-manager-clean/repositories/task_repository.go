package repositories

import (
	"context"
	"task-manager-clean/domain/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepositoryImpl struct {
	collection *mongo.Collection
}

func NewTaskRepository() *TaskRepositoryImpl {
	collection, err := GetTasksCollection()
	if err != nil {
		panic(err)
	}

	return &TaskRepositoryImpl{collection: collection}
}

func (repository *TaskRepositoryImpl) GetByID(id string) (*entities.Task, error) {
	filter := bson.D{{
		Key: "id", Value: id,
	}}

	var task *entities.Task
	err := repository.collection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (repository *TaskRepositoryImpl) GetAll() ([]*entities.Task, error) {
	cur, err := repository.collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	var allTasks []*entities.Task = make([]*entities.Task, 0)
	for cur.Next(context.TODO()) {
		var task *entities.Task
		cur.Decode(&task)

		allTasks = append(allTasks, task)
	}

	return allTasks, nil
}

func (repository *TaskRepositoryImpl) Insert(task *entities.Task) (*entities.Task, error) {
	_, err := repository.collection.InsertOne(context.TODO(), task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (repository *TaskRepositoryImpl) Update(id string, updatedTask *entities.Task) (*entities.Task, error) {
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

	_, err := repository.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	task, err := repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (repository *TaskRepositoryImpl) Delete(id string) (*entities.Task, error) {
	task, err := repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{Key: "id", Value: id}}
	_, err = repository.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	return task, nil
}
