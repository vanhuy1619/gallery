package repositories

import (
	"awesomeProject2/config"
	"awesomeProject2/middleware"
	"awesomeProject2/payload"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type User payload.User

func (User) TableName() string { return "User" }

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

		if user.Username == "" || user.Password == "" {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
			return
		}

		if err := db.Where("username = ?", user.Username).First(&user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// Username doesn't exist in the database
				// Proceed with registration logic
			} else {
				// Error occurred while querying the database
				context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query the database"})
				return
			}
		} else {
			// Username already exists in the database
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Username already exists"})
			return
		}

		hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash password"})
			return
		}

		userValid := User{
			Username: user.Username,
			Password: string(hashed),
		}

		if err := db.Create(&userValid).Error; err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Create new user failed"})
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
