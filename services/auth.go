package services

import (
	"linkmatch-be/database/models"
	"linkmatch-be/requests"

	"github.com/gin-gonic/gin"
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
