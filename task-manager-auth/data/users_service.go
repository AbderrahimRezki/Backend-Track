package data

import (
	"context"
	"errors"
	"task-manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

var ErrUserExists = errors.New("user exists")
var ErrInvalidCredentials = errors.New("invalid username or password")

func GetUserByUsername(username string) (*models.User, error) {
	collection, err := GetUsersCollection()
	if err != nil {
		return nil, err
	}

	filter := bson.D{{
		Key: "username", Value: username,
	}}

	var user *models.User
	err = collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func Login(user *models.User) (*models.User, error) {
	existingUser, err := GetUserByUsername(user.Username)
	if err != nil {
		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)) != nil {
		return nil, ErrInvalidCredentials
	}

	return existingUser, nil
}

func Register(user *models.User) (*models.User, error) {
	if user, err := GetUserByUsername(user.Username); err == nil && user != nil {
		return nil, ErrUserExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)

	collection, err := GetUsersCollection()
	if err != nil {
		return nil, err
	}

	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
