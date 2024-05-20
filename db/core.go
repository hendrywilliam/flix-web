package db

import (
	"fmt"
	"os"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	once       sync.Once
	dbInstance *sqlx.DB
)

func NewDatabase() *sqlx.DB {
	once.Do(func() {
		db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		))
		if err != nil {
			panic(err)
		}
		dbInstance = db
	})

	return dbInstance
}
