package controllers

import (
	"net/http"
	"os"
	"task-manager-clean/domain/entities"
	"task-manager-clean/infrastructure"
	"task-manager-clean/repositories"
	"task-manager-clean/usecases"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	repository *repositories.UserRepositoryImpl
}

func NewAuthController() *AuthController {
	userRepository := repositories.NewUserRepository()
	return &AuthController{repository: userRepository}
}

func (controller *AuthController) Register(ctx *gin.Context) {
	var user entities.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := usecases.Register(controller.repository, &user)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": "registered successfully"})
}

func (controller *AuthController) Login(ctx *gin.Context) {
	var user *entities.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loggedUser, err := usecases.Login(controller.repository, user)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwtToken, err := infrastructure.CreateJWT(loggedUser, os.Getenv("jwt_secret"))
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "description": "error from jwt token creation"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "logged in successfully", "token": jwtToken})
}
