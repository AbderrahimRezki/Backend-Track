package data

import (
	"errors"
	"task-manager/models"
	"time"
)

var tasks = []models.Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

var ErrTaskNotFound = errors.New("task not found")

func GetAllTasks() []models.Task {
	return tasks
}

func GetTask(id string) (*models.Task, error) {
	for _, task := range tasks {
		if task.ID == id {
			return &task, nil
		}
	}

	return nil, ErrTaskNotFound
}

func UpdateTask(id string, updatedTask *models.Task) (*models.Task, error) {
	for idx, task := range tasks {
		if task.ID == id {
			if updatedTask.Title != "" {
				tasks[idx].Title = updatedTask.Title
			}

			if updatedTask.Description != "" {
				tasks[idx].Description = updatedTask.Description
			}

			if updatedTask.Status != "" {
				tasks[idx].Status = updatedTask.Status
			}

			return &tasks[idx], nil
		}
	}

	return nil, ErrTaskNotFound
}

func DeleteTask(id string) (*models.Task, error) {
	for idx, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:idx], tasks[idx+1:]...)
			return &task, nil
		}
	}

	return nil, ErrTaskNotFound
}

func PostTask(task *models.Task) (*models.Task, error) {
	tasks = append(tasks, *task)
	return task, nil
}
