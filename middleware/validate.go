package middleware

import (
	"Consure-App/sdk/response"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWToken(id uint) (string, error) {
	env := os.Getenv("TOKEN_KEY")
	if env == "" {
		return "", fmt.Errorf("error when get env")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
	})
	signedToken, err := token.SignedString([]byte(env))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ValidateJWToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		if bearerToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.ResponseWhenFail("Unauthorized", nil))
			return
		}
		bearerToken = strings.ReplaceAll(bearerToken, "Bearer ", "")
		token, err := jwt.Parse(bearerToken, ekstractToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, response.ResponseWhenFail("Failed to extract token", err.Error()))
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userId := uint(claims["id"].(float64))
			c.Set("id", userId)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusForbidden, response.ResponseWhenFail("Failed to extract token", err.Error()))
			return
		}
	}
}

func ekstractToken(token *jwt.Token) (interface{}, error) {
	env := os.Getenv("TOKEN_KEY")
	if env == "" {
		return "", fmt.Errorf("error when get env")
	}
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, jwt.ErrSignatureInvalid
	}
	return []byte(env), nil
}
