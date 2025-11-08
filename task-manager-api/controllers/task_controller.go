package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"task-manager/data"
	"task-manager/models"
)

func GetAllTasks(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, data.GetAllTasks())
}

func GetTask(ctx *gin.Context) {
	id := ctx.Param("id")
	if task, err := data.GetTask(id); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
	} else {
		ctx.IndentedJSON(http.StatusOK, task)
	}
}

func UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedTask models.Task
	if err := ctx.BindJSON(&updatedTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if task, err := data.UpdateTask(id, &updatedTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})

	} else {
		ctx.IndentedJSON(http.StatusOK, task)
	}
}

func DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")

	if task, err := data.DeleteTask(id); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		ctx.IndentedJSON(http.StatusOK, task)
	}
}

func PostTask(ctx *gin.Context) {
	var newTask models.Task
	if err := ctx.BindJSON(&newTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if task, err := data.PostTask(&newTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		ctx.IndentedJSON(http.StatusOK, task)
	}
}
