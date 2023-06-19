package middleware

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateToken(username string) string {
	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(60 * time.Minute).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY"))) // Replace with your secret key
	if err != nil {
		return ""
	}

	return tokenString
}

var sessions = map[string]session{}

func ValidSession(w http.ResponseWriter, username string) {
	// Session
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	// Set the token in the session map, along with the session information
	sessions[sessionToken] = session{
		Username: username,
		Expiry:   expiresAt,
	}

	// Set the session token as a cookie
	cookie := &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiresAt,
	}
	http.SetCookie(w, cookie)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the token from the request header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is missing"})
			c.Abort()
			return
		}

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

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			c.Abort()
			return
		}

		username, ok := claims["username"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			c.Abort()
			return
		}

		c.Set("username", username)
		fmt.Println(c.GetString("username")) // Print the set username

		//sessionCookie, err := c.Cookie("session_token")
		//if err != nil {
		//	if err == http.ErrNoCookie {
		//		// If the cookie is not set, return an unauthorized status
		//		c.AbortWithStatus(http.StatusUnauthorized)
		//		return
		//	}
		//	// For any other type of error, return a bad request status
		//	c.AbortWithStatus(http.StatusBadRequest)
		//	return
		//}
		//sessionToken := sessionCookie
		//
		//// We then get the session from our session map
		//userSession, exists := sessions[sessionToken]
		//if !exists {
		//	// If the session token is not present in session map, return an unauthorized error
		//	c.AbortWithStatus(http.StatusUnauthorized)
		//	return
		//}
		//// If the session is present, but has expired, we can delete the session, and return
		//// an unauthorized status
		//if userSession.IsExpired() {
		//	delete(sessions, sessionToken)
		//	c.AbortWithStatus(http.StatusUnauthorized)
		//	return
		//}

		// Continue to the next handler
		c.Next()
	}
}

type session struct {
	Username string    `json:"username"`
	Expiry   time.Time `json:"expiry"`
}

func (s session) IsExpired() bool {
	return s.Expiry.Before(time.Now())
}
