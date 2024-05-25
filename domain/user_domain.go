package domain

import (
	"flix/db"

	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/client"
	"gorm.io/gorm"
)

// Form-data or json
type SignupRequest struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type SigninRequest struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type UserRepository interface {
	Create(user *db.User) *gorm.DB
	GetUserByEmail(user *db.User) *gorm.DB
	CreateStripeCustomer(params *stripe.CustomerParams, client *client.API) (*stripe.Customer, error)
	Delete(user *db.User) *gorm.DB
	Update(user *db.User, column string, value any) *gorm.DB
}

type UserUsecase interface {
	Create(user *db.User) *gorm.DB
	GetUserByEmail(user *db.User) *gorm.DB
	CreateStripeCustomer(user *db.User) (*stripe.Customer, error)
	Delete(user *db.User) *gorm.DB
	Update(user *db.User, column string, value any) *gorm.DB
}
