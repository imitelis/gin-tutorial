package middlewares

import (
	"gin-tutorial/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")

		// Check if the Authorization header is present and starts with "Bearer "
		if len(authHeader) <= len(BEARER_SCHEMA) || authHeader[:len(BEARER_SCHEMA)] != BEARER_SCHEMA {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Extract the token string
		tokenString := authHeader[len(BEARER_SCHEMA):] // Get the substring after "Bearer "

		// Validate the token
		token, err := service.NewJWTService().ValidateToken(tokenString)
		if err != nil || !token.Valid {
			log.Println("Invalid token:", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Access the claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			log.Println("Claims[Name]: ", claims["name"])
			log.Println("Claims[Admin]: ", claims["admin"])
			log.Println("Claims[Issuer]: ", claims["iss"])
			log.Println("Claims[IssuedAt]: ", claims["iat"])
			log.Println("Claims[ExpiresAt]: ", claims["exp"])
		}

		// Proceed to the next handler
		c.Next()
	}
}
