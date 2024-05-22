package db

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	once       sync.Once
	dbInstance *gorm.DB
)

func NewDatabase() *gorm.DB {
	once.Do(func() {
		dbCredentials := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
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
		start := time.Now()
		db.AutoMigrate(&User{}, &Profile{}, &WatchingList{})
		elapsed := time.Since(start)
		log.Printf("Migration succeeded. Time elapsed: %s", elapsed)

		dbInstance = db
	})

	return dbInstance
}
