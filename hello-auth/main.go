package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"hello-auth/models"
)

var jwtSecret = []byte("MySecret")

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if strings.Trim(authHeader, " ") == "" {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "auth header required"})
			ctx.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid auth header"})
			ctx.Abort()
			return
		}

		token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func main() {
	users := make(map[string]*models.User)
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "OK"})
	})

	router.POST("/register", func(ctx *gin.Context) {
		var user models.User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user.Password = string(hashedPwd)
		users[user.Email] = &user
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "user registered"})
	})

	router.POST("/login", func(ctx *gin.Context) {
		var user models.User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		existingUser, ok := users[user.Email]
		if !ok || bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)) != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": user.ID,
			"email":   user.Email,
		})

		jwtToken, err := token.SignedString(jwtSecret)
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "user successfully logged in", "token": jwtToken})
	})

	router.GET("/secure", AuthMiddleWare(), func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "This is a secure route"})

	})

	router.Run("localhost:8080")
}
