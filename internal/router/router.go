package router

import (
	"github.com/IgorBrizack/taghos-backend/internal/controller"
	"github.com/IgorBrizack/taghos-backend/internal/database"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *database.Database, bookController *controller.BookController) *gin.Engine {
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

	r.POST("/books", bookController.CreateBook)
	r.GET("/books/:id", bookController.GetBook)
	r.GET("/books", bookController.GetAllBooks)
	r.DELETE("/books/:id", bookController.DeleteBook)
	r.PUT("/books/:id", bookController.UpdateBook)

	return r
}
