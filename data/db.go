package data

import (
	"compass-wealth/enums"
	"compass-wealth/model"
	"compass-wealth/utils"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/log"

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
	err := DB.AutoMigrate(
		&model.Account{},
		&model.Asset{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	fmt.Println("Database migration completed successfully")
}

func SeedDB() {
	adminEmail := os.Getenv("AdminEmail")

	if strings.TrimSpace(adminEmail) == "" {
		log.Fatal("Admin email not found")
	}

	adminName := os.Getenv("AdminName")
	if strings.TrimSpace(adminName) == "" {
		log.Fatal("Admin name not found")
	}

	var admin model.Account

	err := DB.Where("email = ?", adminEmail).First(&admin).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hashedPassword, err := utils.HashPassword("QWRtaW5AMTIz")

			if err != nil {
				log.Fatalf("Failed to hash password: %v", err)
			}

			admin = model.Account{
				UserName: adminName,
				Email:    adminEmail,
				Password: hashedPassword,
				Role:     enums.Admin,
			}

			if err := DB.Create(&admin).Error; err != nil {
				log.Printf("Failed to seed admin user: %v", err)
			} else {
				fmt.Println("Admin user seeded successfully")
			}
		} else {
			log.Printf("Error checking for existing admin: %v", err)
		}
	}
}
