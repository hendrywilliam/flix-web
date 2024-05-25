package repository

import (
	"flix/db"

	"gorm.io/gorm"
)

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{database: db}
}

func (ur *userRepository) Create(user *db.User) *gorm.DB {
	result := ur.database.Create(user)
	return result
}
