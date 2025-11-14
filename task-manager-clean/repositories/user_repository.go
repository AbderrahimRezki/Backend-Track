package repositories

import (
	"context"
	"errors"
	"task-manager-clean/domain/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var ErrNotImplemented = errors.New("not implemented")

type UserRepositoryImpl struct {
	collection *mongo.Collection
}

func NewUserRepository() *UserRepositoryImpl {
	collection, err := GetUsersCollection()
	if err != nil {
		panic(err)
	}
	return &UserRepositoryImpl{collection: collection}
}

func (repository *UserRepositoryImpl) GetByID(username string) (*entities.User, error) {
	filter := bson.D{{
		Key: "username", Value: username,
	}}

	var user *entities.User
	err := repository.collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) GetAll() ([]*entities.User, error) {
	filter := bson.D{{}}

	cur, err := repository.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var users []*entities.User = make([]*entities.User, 0)
	for cur.Next(context.TODO()) {
		var user *entities.User
		cur.Decode(&user)
		users = append(users, user)
	}

	return users, nil
}

func (repository *UserRepositoryImpl) Insert(user *entities.User) (*entities.User, error) {
	_, err := repository.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}

	return repository.GetByID(user.Username)
}

func (repository *UserRepositoryImpl) Update(username string, updatedUser *entities.User) (*entities.User, error) {
	return nil, ErrNotImplemented
}
