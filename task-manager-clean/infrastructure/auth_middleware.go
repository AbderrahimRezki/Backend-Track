package infrastructure

import (
	"errors"
	"net/http"
	"slices"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var ErrMissingAuthHeader = errors.New("missing auth header")
var ErrInvalidAuthHeader = errors.New("invalid auth header")
var ErrUnexpectedSigningAlg = errors.New("unexpected signing algorithm")

func getTokenString(ctx *gin.Context) (string, error) {
	authHeader := ctx.GetHeader("Authorization")
	if strings.Trim(authHeader, " ") == "" {
		return "", ErrMissingAuthHeader
	}

	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
		return "", ErrInvalidAuthHeader
	}

	tokenStr := strings.Trim(authParts[1], " ")
	return tokenStr, nil
}

func AuthenticationMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr, err := getTokenString(ctx)
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		token, err := ParseToken(tokenStr, jwtSecret)
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		if err = IsTokenExpired(token); err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		ctx.Set("token", token)
		ctx.Next()
	}
}

func AuthorizationMiddleware(allowedRoles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, ok := ctx.MustGet("token").(*jwt.Token)
		if !ok {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to get token"})
			ctx.Abort()
			return
		}

		roles, err := GetRoles(token)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		valid := false
		for _, role := range roles {
			if slices.Contains(allowedRoles, role) {
				valid = true
				break
			}
		}

		if !valid {
			ctx.IndentedJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
