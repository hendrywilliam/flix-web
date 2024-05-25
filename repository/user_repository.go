package repository

import (
	"flix/db"

	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/client"
	"gorm.io/gorm"
)

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{database: db}
}

func (ur *userRepository) Create(user *db.User) *gorm.DB {
	return ur.database.Create(user)
}

func (ur *userRepository) Delete(user *db.User) *gorm.DB {
	return ur.database.Where("email = ?", user.Email).Delete(&user)
}

func (ur *userRepository) Update(user *db.User, column string, value any) *gorm.DB {
	return ur.database.Model(&user).Update(column, value)
}

func (ur *userRepository) GetUserByEmail(user *db.User) *gorm.DB {
	return ur.database.Where("email = ?", user.Email).Find(user)
}

func (ur *userRepository) CreateStripeCustomer(params *stripe.CustomerParams, client *client.API) (*stripe.Customer, error) {
	customer, err := client.Customers.New(params)
	if err != nil {
		return nil, err
	}
	return customer, nil
}
