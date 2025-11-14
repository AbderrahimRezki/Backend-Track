package repositories

import (
	"task-manager-clean/domain/entities"
)

type UserRepository interface {
	GetByID(string) (*entities.User, error)
	GetAll() ([]*entities.User, error)
	Insert(*entities.User) (*entities.User, error)
	Update(string, *entities.User) (*entities.User, error)
}
