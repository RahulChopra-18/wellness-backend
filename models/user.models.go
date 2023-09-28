package models

import "gorm.io/gorm"

// SignInInput struct
type SignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	gorm.Model
	Name       string `gorm:"type:varchar(255);not null"`
	Email      string `gorm:"type:varchar(255);unique_index;not null"`
	Password   string `gorm:"type:varchar(255);not null"`
	Sex        string `gorm:"type:enum('Male','Female','Other');not null"`
	BloodGroup string `gorm:"type:varchar(5);not null"`
	Height     float64
	Weight     float64
	FitnessApp string `gorm:"type:enum('Google Fit','Apple Fit','Samsung Fit')"`
}

type CreateUserParams struct {
	gorm.Model
	Name       string `gorm:"type:varchar(255);not null"`
	Email      string `gorm:"type:varchar(255);unique_index;not null"`
	Password   string `gorm:"type:varchar(255);not null"`
	Sex        string `gorm:"type:enum('Male','Female','Other');not null"`
	BloodGroup string `gorm:"type:varchar(5);not null"`
	Height     float64
	Weight     float64
	FitnessApp string `gorm:"type:enum('Google Fit','Apple Fit','Samsung Fit')"`
}

type UserResponse struct {
	gorm.Model
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Sex        string  `json:"sex"`
	BloodGroup string  `json:"bloodGroup"`
	Height     float64 `json:"height"`
	Weight     float64 `json:"weight"`
	FitnessApp string  `json:"fitnessApp"`
}

func FilteredResponse(user *User) UserResponse {
	return UserResponse{
		Model:      user.Model,
		Name:       user.Name,
		Email:      user.Email,
		Sex:        user.Sex,
		BloodGroup: user.BloodGroup,
		Height:     user.Height,
		Weight:     user.Weight,
		FitnessApp: user.FitnessApp,
	}
}
