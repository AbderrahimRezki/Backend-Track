package routers

import (
	"net/http"
	"os"
	"task-manager-clean/delivery/controllers"
	"task-manager-clean/infrastructure"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	authController := controllers.NewAuthController()
	taskController := controllers.NewTaskController()

	router.GET("/", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Server up!"})
	})

	authRouter := router.Group("/auth")
	{
		authRouter.POST("/register", authController.Register)
		authRouter.POST("/login", authController.Login)
	}

	tasksRouter := router.Group("/tasks",
		infrastructure.AuthenticationMiddleware(os.Getenv("jwt_secret")),
		infrastructure.AuthorizationMiddleware([]string{"user", "admin"}))
	{
		tasksRouter.GET("/", taskController.GetAllTasks)
		tasksRouter.GET("/:id", taskController.GetTask)
		tasksRouter.PUT("/:id", taskController.UpdateTask)
		tasksRouter.DELETE("/:id", taskController.DeleteTask)
		tasksRouter.POST("/", taskController.PostTask)
	}

	adminRouter := router.Group("/admin",
		infrastructure.AuthenticationMiddleware(os.Getenv("jwt_secret")),
		infrastructure.AuthorizationMiddleware([]string{"admin"}))
	{
		adminRouter.GET("/", func(ctx *gin.Context) {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "you proved you are an admin."})
		})
	}

	return router
}
