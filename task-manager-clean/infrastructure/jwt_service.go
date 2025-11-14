package infrastructure

import (
	"errors"
	"strings"
	"task-manager-clean/domain/entities"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var ErrInvalidToken = errors.New("invalid token")
var ErrInvalidClaims = errors.New("invalid claims")
var ErrInvalidTokenExp = errors.New("invalid token expiration")
var ErrTokenExpired = errors.New("token expired")
var ErrMissingRoles = errors.New("roles field is missing from token")

func CreateJWT(user *entities.User, jwtSecret string) (string, error) {
	roles := make([]string, 0)
	roles = append(roles, user.Roles...)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"roles":    roles,
		"exp":      time.Now().Add(3 * time.Minute),
	})

	jwtToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func ParseToken(tokenStr string, jwtSecret string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnexpectedSigningAlg
		}

		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, ErrInvalidToken
	}

	return token, nil
}

func IsTokenExpired(token *jwt.Token) error {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ErrInvalidClaims
	}

	expStr, ok := claims["exp"].(string)
	if !ok {
		return ErrInvalidTokenExp
	}

	exp, err := time.Parse(time.RFC3339, expStr)
	if err != nil {
		return err
	}

	if time.Now().After(exp) {
		return ErrTokenExpired
	}

	return nil
}

func GetRoles(token *jwt.Token) ([]string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrInvalidClaims
	}

	rawRoles, ok := claims["roles"].([]interface{})
	if !ok {
		return nil, ErrMissingRoles
	}

	roles := make([]string, 0)
	for _, role := range rawRoles {
		if roleStr, ok := role.(string); ok && strings.Trim(roleStr, " ") != "" {
			roles = append(roles, roleStr)
		}
	}

	return roles, nil
}
