package controllers

import (
	"net/http"
	"task-manager-clean/domain/entities"
	"task-manager-clean/repositories"
	"task-manager-clean/usecases"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	repository *repositories.TaskRepositoryImpl
}

func NewTaskController() *TaskController {
	taskRespository := repositories.NewTaskRepository()
	return &TaskController{repository: taskRespository}
}

func (controller *TaskController) GetAllTasks(ctx *gin.Context) {
	tasks, err := usecases.GetAllTasks(controller.repository)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, tasks)
}

func (controller *TaskController) GetTask(ctx *gin.Context) {
	id := ctx.Param("id")
	if task, err := usecases.GetTask(controller.repository, id); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		ctx.IndentedJSON(http.StatusOK, task)
	}
}

func (controller *TaskController) UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedTask entities.Task
	if err := ctx.BindJSON(&updatedTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if task, err := usecases.UpdateTask(controller.repository, id, &updatedTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})

	} else {
		ctx.IndentedJSON(http.StatusOK, task)
	}
}

func (controller *TaskController) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")

	if task, err := usecases.DeleteTask(controller.repository, id); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		ctx.IndentedJSON(http.StatusOK, task)
	}
}

func (controller *TaskController) PostTask(ctx *gin.Context) {
	var newTask entities.Task
	if err := ctx.BindJSON(&newTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if task, err := usecases.AddTask(controller.repository, &newTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		ctx.IndentedJSON(http.StatusOK, task)
	}
}
