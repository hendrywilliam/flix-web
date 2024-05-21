package db

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	once       sync.Once
	dbInstance *gorm.DB
)

func NewDatabase() *gorm.DB {
	once.Do(func() {
		dbCredentials := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)
		db, err := gorm.Open(postgres.Open(dbCredentials), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		dbInstance = db
	})
	return dbInstance
}
