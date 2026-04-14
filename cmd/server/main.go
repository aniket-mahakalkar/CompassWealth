package main

import (
	"compass-wealth/data"
	"compass-wealth/web"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Compass Wealth")

	data.InitDB()
	data.MigrateDB()
	
	fmt.Println("Server is starting...")
	web.Init()
}
