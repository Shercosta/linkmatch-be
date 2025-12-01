package services

import (
	"linkmatch-be/database/models"

	"gorm.io/gorm"
)

func GetUser(db *gorm.DB, username string) *models.User {
	var user models.User
	result := db.Where("username = ?", username).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return nil
	}

	return &user
}

func GetUserPublic(db *gorm.DB, username string) *models.UserPublic {
	var user models.UserPublic
	result := db.
		Preload("Image").
		Where("username = ?", username).
		First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return nil
	}

	return &user
}

func GetUsers(db *gorm.DB) []models.UserPublic {
	var users []models.UserPublic

	result := db.Find(&users)

	if result.Error != nil {
		return []models.UserPublic{}
	}

	return users
}
