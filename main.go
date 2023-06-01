package main

import (
	"fmt"
	"structure/router"

	"github.com/joho/godotenv"
)

//

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("env not found")
		return
	}
	router.Routing()
}
