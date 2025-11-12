package controllers

import (
	"fmt"
	"net/http"
	"os"
	"task-manager/data"
	"task-manager/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateJWT(user *models.User) (string, error) {
	roles := make([]string, 0)
	roles = append(roles, user.Roles...)
	fmt.Println(roles)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"roles":    roles,
		"exp":      time.Now().Add(3 * time.Minute),
	})

	jwtToken, err := token.SignedString([]byte(os.Getenv("jwt_secret")))
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func Register(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := data.Register(&user)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": "registered successfully"})
}

func Login(ctx *gin.Context) {
	var user *models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loggedUser, err := data.Login(user)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwtToken, err := CreateJWT(loggedUser)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "description": "error from jwt token creation"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "logged in successfully", "token": jwtToken})
}
