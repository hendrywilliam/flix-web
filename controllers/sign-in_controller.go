package controllers

import (
	"flix/domain"
	// "github.com/gin-gonic/gin"
)

type SigninController struct {
	userUsecase domain.UserUsecase
}

// func (si *SigninController) Signin(c *gin.Context) {
// 	var request domain.SigninRequest
// }
