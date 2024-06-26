package main

import (
	"flix/db"
	"flix/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	env_err := godotenv.Load(".env")
	if env_err != nil {
		panic(env_err)
	}

	mode := os.Getenv("GIN_MODE")
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	db := db.NewDatabase()

	defer func() {
		// Clean up
		db, _ := db.DB()
		db.Close()
	}()
	r := routes.RoutesCore(db)

	address := "localhost:8080"
	fmt.Printf("Server started at %s \n", address)
	r.Run(address)
}
