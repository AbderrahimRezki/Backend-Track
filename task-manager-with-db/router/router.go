package router

import (
	"net/http"
	"task-manager/controllers"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Server up!"})
	})

	tasksRouter := router.Group("/tasks")
	{
		tasksRouter.GET("/", controllers.GetAllTasks)
		tasksRouter.GET("/:id", controllers.GetTask)
		tasksRouter.PUT("/:id", controllers.UpdateTask)
		tasksRouter.DELETE("/:id", controllers.DeleteTask)
		tasksRouter.POST("/", controllers.PostTask)
	}

	return router
}
