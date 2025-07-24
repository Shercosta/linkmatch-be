package services

import (
	"linkmatch-be/database/models"
	"linkmatch-be/requests"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(db *gorm.DB, c *gin.Context, object *requests.RegisterRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(object.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := &models.User{
		Username: object.Username,
		Password: string(hashedPassword),
	}

	if err := db.Create(newUser).Error; err != nil {
		return err
	}

	return nil
}

func Login(db *gorm.DB, object *requests.RegisterRequest, user *models.User) (obj any, err error) {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(object.Password)); err != nil {
		return nil, err
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	constructResponse := map[string]string{
		"username": user.Username,
		"token":    tokenString,
	}

	return constructResponse, nil
}
