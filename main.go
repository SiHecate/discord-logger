package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Welcome to discord logger bot")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

}
