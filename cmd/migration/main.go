package main

import (
	"compass-wealth/data"
	"compass-wealth/model"

	"github.com/charmbracelet/log"
)

func main() {
	log.Info("Running database migrations...")

	data.InitDB()
	data.MigrateDB()

	log.Info("Migration command finished successfully")
}
