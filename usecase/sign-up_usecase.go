package usecase

import (
	"flix/db"
	"flix/domain"

	"gorm.io/gorm"
)

type signupUsecase struct {
	signupRepository domain.SignupRepository
}

func NewSignupUsecase(signupRepository domain.SignupRepository) domain.SignupUsecase {
	return &signupUsecase{
		signupRepository: signupRepository}
}

func (su *signupUsecase) Create(user *db.User) *gorm.DB {
	return su.signupRepository.Create(user)
}
