package usecase

import (
	"flix/db"
	"flix/domain"
	"flix/internal/libs"
	"fmt"

	"github.com/stripe/stripe-go/v78"
	"gorm.io/gorm"
)

type signupUsecase struct {
	userRepository domain.UserRepository
}

func NewSignupUsecase(userRepository domain.UserRepository) domain.UserUsecase {
	return &signupUsecase{
		userRepository: userRepository}
}

func (su *signupUsecase) Create(user *db.User) *gorm.DB {
	return su.userRepository.Create(user)
}

func (su *signupUsecase) Delete(user *db.User) *gorm.DB {
	return su.userRepository.Delete(user)
}

func (su *signupUsecase) Update(user *db.User, column string, value any) *gorm.DB {
	return su.userRepository.Update(user, column, value)
}

func (su *signupUsecase) GetUserByEmail(user *db.User) *gorm.DB {
	return su.userRepository.GetUserByEmail(user)
}

func (su *signupUsecase) CreateStripeCustomer(user *db.User) (*stripe.Customer, error) {
	stripeClient := libs.StripeClient()
	description := fmt.Sprintf("Customer with email: %s", user.Email)
	params := &stripe.CustomerParams{
		Description: &description,
		Email:       &user.Email,
	}
	return su.userRepository.CreateStripeCustomer(params, stripeClient)
}
