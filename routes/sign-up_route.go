package routes

import (
	"flix/controllers"
	"flix/repository"
	"flix/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewSignupRoute(r *gin.RouterGroup, db *gorm.DB) {
	ur := repository.NewUserRepository(db)
	uc := &controllers.SignupController{
		UserUsecase: usecase.NewSignupUsecase(ur),
	}

	r.POST("/sign-up", uc.Signup)
}
