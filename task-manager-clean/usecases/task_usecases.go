package usecases

import (
	"task-manager-clean/domain/entities"
	"task-manager-clean/domain/repositories"
)

func GetTask(repository repositories.TaskRepository, id string) (*entities.Task, error) {
	return repository.GetByID(id)
}

func GetAllTasks(repository repositories.TaskRepository) ([]*entities.Task, error) {
	return repository.GetAll()
}

func AddTask(repository repositories.TaskRepository, task *entities.Task) (*entities.Task, error) {
	return repository.Insert(task)
}

func UpdateTask(repository repositories.TaskRepository, id string, updatedTask *entities.Task) (*entities.Task, error) {
	return repository.Update(id, updatedTask)
}

func DeleteTask(repository repositories.TaskRepository, id string) (*entities.Task, error) {
	return repository.Delete(id)
}
