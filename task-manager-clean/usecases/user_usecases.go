package usecases

import (
	"errors"
	"task-manager-clean/domain/entities"
	"task-manager-clean/domain/repositories"
	"task-manager-clean/infrastructure"
)

var ErrUserExists = errors.New("user exists")
var ErrInvalidCredentials = errors.New("invalid credentials")

func GetUser(repository repositories.UserRepository, username string) (*entities.User, error) {
	return repository.GetByID(username)
}

func Register(repository repositories.UserRepository, user *entities.User) (*entities.User, error) {
	if user, err := GetUser(repository, user.Username); err == nil && user != nil {
		return nil, ErrUserExists
	}

	hashedPassword, err := infrastructure.GetHashedPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword
	return repository.Insert(user)
}

func Login(repository repositories.UserRepository, user *entities.User) (*entities.User, error) {
	existingUser, err := GetUser(repository, user.Username)
	if err != nil {
		return nil, err
	}

	if infrastructure.CompareHashAndPassword(existingUser.Password, user.Password) {
		return existingUser, nil
	}

	return nil, ErrInvalidCredentials
}
