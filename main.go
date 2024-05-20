package main

import (
	"flix/routes"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	env_err := godotenv.Load(".env")
	if env_err != nil {
		panic(env_err)
	}

	r := routes.RoutesCore()

	address := "localhost:8080"
	fmt.Printf("Server started at %s \n", address)
	
	errServe := http.ListenAndServe(address, r)
	if errServe != nil {
		panic(errServe)
	}
}