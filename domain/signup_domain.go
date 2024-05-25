package domain

import (
	"flix/db"

	"gorm.io/gorm"
)

// Form-data or json
type SignupRequest struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type SignupRepository interface {
	Create(user *db.User) *gorm.DB
}

type SignupUsecase interface {
	Create(user *db.User) *gorm.DB
}
