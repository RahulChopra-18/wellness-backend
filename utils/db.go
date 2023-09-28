package utils

import (
	"Wellness-monitoring/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUserByID(db *gorm.DB, userID uuid.UUID) (*models.User, error) {
	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(db *gorm.DB, email string) (*models.User, error) {
	var user models.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
