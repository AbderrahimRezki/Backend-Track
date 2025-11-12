package router

import (
	"net/http"
	"task-manager/controllers"
	"task-manager/middleware"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Server up!"})
	})

	authRouter := router.Group("/auth")
	{
		authRouter.POST("/register", controllers.Register)
		authRouter.POST("/login", controllers.Login)
	}

	tasksRouter := router.Group("/tasks",
		middleware.AuthenticationMiddleware(),
		middleware.TokenExpirationMiddleware(),
		middleware.AuthorizationMiddleware([]string{"user", "admin"}))
	{
		tasksRouter.GET("/", controllers.GetAllTasks)
		tasksRouter.GET("/:id", controllers.GetTask)
		tasksRouter.PUT("/:id", controllers.UpdateTask)
		tasksRouter.DELETE("/:id", controllers.DeleteTask)
		tasksRouter.POST("/", controllers.PostTask)
	}

	adminRouter := router.Group("/admin",
		middleware.AuthenticationMiddleware(),
		middleware.TokenExpirationMiddleware(),
		middleware.AuthorizationMiddleware([]string{"admin"}))
	{
		adminRouter.GET("/", func(ctx *gin.Context) {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "you proved you are an admin."})
		})
	}

	return router
}
