package controllers

import (
	"flix/db"
	"flix/domain"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
}

// Create user
func (su *SignupController) Signup(c *gin.Context) {
	var request domain.SignupRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	user := db.User{Email: request.Email, Password: request.Password}

	// Create hash
	hash, err := argon2id.CreateHash(user.Password, argon2id.DefaultParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "Internal Server Error.",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	// Change plain text pass to hash.
	user.Password = hash

	result := su.SignupUsecase.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "Failed to create user.",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	if result.RowsAffected > 0 {
		c.JSON(http.StatusCreated, domain.SuccessResponse[string]{
			Message: "User created.",
			Code:    http.StatusCreated,
			Data:    user.Email,
		})
		return
	}
}
