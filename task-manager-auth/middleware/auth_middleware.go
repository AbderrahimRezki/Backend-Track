package middleware

import (
	"errors"
	"net/http"
	"os"
	"slices"
	"strings"
	"time"

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

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr, err := getTokenString(ctx)
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, ErrUnexpectedSigningAlg
			}

			return []byte(os.Getenv("jwt_secret")), nil
		})

		if err != nil || !token.Valid {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		ctx.Set("token", token)
		ctx.Next()
	}
}

func TokenExpirationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, ok := ctx.MustGet("token").(*jwt.Token)
		if !ok {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to get token"})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid claims"})
			ctx.Abort()
			return
		}

		expStr, ok := claims["exp"].(string)
		if !ok {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "invalid token expiration date"})
			ctx.Abort()
			return
		}
		exp, err := time.Parse(time.RFC3339, expStr)
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "invalid token expiration date"})
		}
		if time.Now().After(exp) {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
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

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid claims"})
			ctx.Abort()
			return
		}

		rawRoles, ok := claims["roles"].([]interface{})
		if !ok {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "roles missing from claims"})
			ctx.Abort()
			return
		}

		roles := make([]string, 0)
		for _, role := range rawRoles {
			if roleStr, ok := role.(string); ok && strings.Trim(roleStr, " ") != "" {
				roles = append(roles, roleStr)
			}
		}

		if len(roles) == 0 {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "roles missing from claims"})
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
	}
}
