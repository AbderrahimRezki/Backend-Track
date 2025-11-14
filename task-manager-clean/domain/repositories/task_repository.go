package repositories

import "task-manager-clean/domain/entities"

type TaskRepository interface {
	GetByID(string) (*entities.Task, error)
	GetAll() ([]*entities.Task, error)
	Insert(*entities.Task) (*entities.Task, error)
	Update(string, *entities.Task) (*entities.Task, error)
	Delete(string) (*entities.Task, error)
}
