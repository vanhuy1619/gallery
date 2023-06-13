package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the token from the request header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is missing"})
			c.Abort()
			return
		}

		//// Extract the token string from the header
		//splitToken := strings.Split(authHeader, "Bearer ")
		//if len(splitToken) != 2 {
		//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token format"})
		//	c.Abort()
		//	return
		//}
		//tokenString := splitToken[1]

		// Validate and parse the token
		token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
			// Replace "your-secret-key" with your actual secret key used for token signing
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			c.Abort()
			return
		}

		claims, _ := token.Claims.(jwt.MapClaims)
		username := claims["username"].(string)
		c.Set("username", username)
		fmt.Print(c.Get("username")) //print set username
		// Continue to the next handler
		c.Next()
	}
}
