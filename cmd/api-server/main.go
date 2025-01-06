package main

import (
	"log"

	"github.com/IgorBrizack/taghos-backend/internal/controller"
	"github.com/IgorBrizack/taghos-backend/internal/database"
	migrations "github.com/IgorBrizack/taghos-backend/internal/database/migration"
	"github.com/IgorBrizack/taghos-backend/internal/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.Close()

	migrations.ApplyMigrations(db.GetConnection())

	bookRepository := repository.NewBookRepository(db.GetConnection())

	r := gin.Default()

	r.GET("/health-check", func(c *gin.Context) {
		if err := db.ValidateConnection(); err != nil {
			c.JSON(500, gin.H{
				"message": "Database connection failed",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Running",
		})
	})

	log.Println("Starting API on port 8150")

	bookController := controller.NewBookController(bookRepository)

	r.POST("/books", bookController.CreateBook)
	r.GET("/books/:id", bookController.GetBook)
	r.GET("/books", bookController.GetAllBooks)
	r.DELETE("/books/:id", bookController.DeleteBook)
	r.PUT("/books/:id", bookController.UpdateBook)

	if err := r.Run(":8150"); err != nil {
		log.Fatalf("Error starting API: %v", err)
	}
}
