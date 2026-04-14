package data

import (
	"compass-wealth/model"
	"fmt"
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("./compass_wealth.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	
	fmt.Println("Successfully connected to SQLite database using GORM (modernc.org/sqlite)")
}

func MigrateDB() {
	if DB == nil {
		log.Fatal("Database not initialized. Call InitDB first.")
	}
	err := DB.AutoMigrate(&model.Account{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	fmt.Println("Database migration completed successfully")
}
