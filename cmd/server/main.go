package main

import (
	"compass-wealth/data"
	"compass-wealth/web"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	fmt.Println("Welcome to Compass Wealth")

	data.InitDB()
	data.MigrateDB()
	data.SeedDB()
	
	fmt.Println("Server is starting...")
	web.Init()
}
