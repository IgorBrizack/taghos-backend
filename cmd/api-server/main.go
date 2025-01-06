package main

import (
	"log"

	"github.com/IgorBrizack/taghos-backend/internal/controller"
	"github.com/IgorBrizack/taghos-backend/internal/database"
	migrations "github.com/IgorBrizack/taghos-backend/internal/database/migration"
	"github.com/IgorBrizack/taghos-backend/internal/repository"
	"github.com/IgorBrizack/taghos-backend/internal/router"
)

func main() {
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.Close()

	migrations.ApplyMigrations(db.GetConnection())

	bookRepository := repository.NewBookRepository(db.GetConnection())

	bookController := controller.NewBookController(bookRepository)

	r := router.SetupRouter(db, bookController)

	log.Println("Starting API on port 8150")
	if err := r.Run(":8150"); err != nil {
		log.Fatalf("Error starting API: %v", err)
	}
}
