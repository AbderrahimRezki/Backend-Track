package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

var tasks = []Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

func main() {
	router := gin.Default()
	router.GET("/ping", pong)
	router.GET("/tasks", getAllTasks)
	router.GET("/tasks/:id", getTask)
	router.PUT("/tasks/:id", updateTask)
	router.DELETE("/tasks/:id", deleteTask)
	router.POST("/tasks", postTask)
	router.Run("localhost:8080")
}

func pong(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "All good"})
}

func getAllTasks(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, tasks)
}

func getTask(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, task := range tasks {
		if task.ID == id {
			ctx.IndentedJSON(http.StatusOK, task)
			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task Not Found."})
}

func updateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedTask Task

	if err := ctx.BindJSON(&updatedTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	for idx, task := range tasks {
		if task.ID == id {
			if updatedTask.Title != "" {
				tasks[idx].Title = updatedTask.Title
			}

			if updatedTask.Description != "" {
				tasks[idx].Description = updatedTask.Description
			}

			ctx.IndentedJSON(http.StatusOK, tasks[idx])
			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found."})
}

func deleteTask(ctx *gin.Context) {
	id := ctx.Param("id")

	for idx, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:idx], tasks[idx+1:]...)
			ctx.IndentedJSON(http.StatusOK, task)
			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task Not Found."})
}

func postTask(ctx *gin.Context) {
	var newTask Task

	if err := ctx.BindJSON(&newTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	for _, task := range tasks {
		if task.ID == newTask.ID {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Task with id already exists"})
			return
		}
	}

	tasks = append(tasks, newTask)
	ctx.IndentedJSON(http.StatusCreated, newTask)
}
