package main

import (
	"fmt"
	"structure/database"
	"structure/router"

	"github.com/joho/godotenv"
)

// init
func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("env not found")
		return
	}
	client, _, err := database.ConnectDb()
	if err != nil {
		panic("db is not running")
	}
	database.Client = client
}

func main() {

	router.Routing()
}
