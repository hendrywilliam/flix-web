package db

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email                string `gorm:"unique;column:email"`
	Password             string
	StripeSubscriptionID string    `gorm:"unique;column:stripe_subscription_id"`
	StripeCustomerID     string    `gorm:"unique;column:stripe_customer_id"`
	Profile              []Profile `gorm:"foreignKey:UserID"`
}

type Profile struct {
	gorm.Model
	UserID       int `gorm:"column:user_id"`
	Name         string
	Avatar       string
	WatchingList []WatchingList `gorm:"foreignKey:ProfileId"`
}

type WatchingList struct {
	gorm.Model
	ProfileId int `gorm:"column:profile_id"`
	Movie     datatypes.JSON
}
