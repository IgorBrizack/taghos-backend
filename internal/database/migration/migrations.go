package migrations

import (
	"log"

	"github.com/IgorBrizack/taghos-backend/internal/models"
	"gorm.io/gorm"
)

func ApplyMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Book{},
	)
	if err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}
	log.Println("Migrations applied successfully.")
}
