package controller

import (
	"net/http"
	"strconv"

	"github.com/IgorBrizack/taghos-backend/internal/command"
	"github.com/IgorBrizack/taghos-backend/internal/repository"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookRepository repository.BookRepository
}

func NewBookController(bookRepository repository.BookRepository) *BookController {
	return &BookController{bookRepository: bookRepository}
}

func (ctrl *BookController) CreateBook(c *gin.Context) {
	var params command.CreateBookParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createBookCmd := command.NewCreateBookCommand(ctrl.bookRepository)

	book, err := createBookCmd.Execute(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": book})
}

func (ctrl *BookController) GetBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	getBookCmd := command.NewGetBookCommand(ctrl.bookRepository)
	book, err := getBookCmd.Execute(command.GetBookParams{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if book == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (ctrl *BookController) GetAllBooks(c *gin.Context) {
	getAllBooksCmd := command.NewGetAllBooksCommand(ctrl.bookRepository)
	books, err := getAllBooksCmd.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if books == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No books found"})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (ctrl *BookController) DeleteBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	deleteBookCmd := command.NewDeleteBookCommand(ctrl.bookRepository)
	err = deleteBookCmd.Execute(command.DeleteBookParams{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

func (ctrl *BookController) UpdateBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var params command.UpdateBookParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	params.ID = id

	updateBookCmd := command.NewUpdateBookCommand(ctrl.bookRepository)

	updatedBook, err := updateBookCmd.Execute(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedBook)
}
