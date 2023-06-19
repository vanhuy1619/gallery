package repositories

import (
	"awesomeProject2/config"
	"awesomeProject2/middleware"
	"awesomeProject2/model"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type User model.User

func (User) TableName() string { return "users" }

func init() {
	log.Println("Main initialization, load config file")
	config.LoadConfig()
}

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var user User
		if err := context.ShouldBindJSON(&user); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		var existingUser User
		if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
				return
			}
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query the database"})
			return
		}

		// Verify the password
		if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)); err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		token := middleware.GenerateToken(existingUser.Username)
		middleware.ValidSession(context.Writer, user.Username)
		context.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func Regist(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var user User
		if err := context.ShouldBindJSON(&user); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		if user.Username == "" || user.Password == "" || user.Email == "" {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
			return
		}

		// Format the JSON string
		jsonData, err := json.MarshalIndent(user, "", "\t")
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to format JSON"})
			return
		}

		// Print the formatted JSON
		fmt.Println(string(jsonData))

		var existingUser User
		if err := db.Where("username = ? OR email = ?", user.Username, user.Email).First(&existingUser).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// Username and email combination doesn't exist in the database
				// Proceed with registration logic
			} else {
				// Error occurred while querying the database
				context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query the database"})
				return
			}
		} else {
			// Username or email already exists in the database
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Username or email already exists"})
			return
		}

		hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash password"})
			return
		}

		_, errE := config.IsEmailValid(user.Email)
		if errE != nil {
			context.JSON(http.StatusBadRequest, model.ResponeError{Code: 1, Message: "Invalid Email"})
			return
		}

		userValid := User{
			Username: user.Username,
			Password: string(hashed),
			Gender:   user.Gender,
			Email:    user.Email,
		}

		if err := db.Create(&userValid).Error; err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create new user"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"username": userValid.Username, "code": "0"})
	}
}

type LoginRequest struct {
	username string `json:"username"`
	password string `json:"password"`
}

type LoginRespone struct {
	Token string `json:"token"`
}
