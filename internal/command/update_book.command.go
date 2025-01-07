package command

import (
	"log"

	"github.com/IgorBrizack/taghos-backend/internal/models"
	"github.com/IgorBrizack/taghos-backend/internal/repository"
)

type UpdateBookParams struct {
	ID       int64  `json:"id"`
	Title    string `json:"title" binding:"required"`
	Category string `json:"category" binding:"required"`
	Author   string `json:"author" binding:"required"`
	Synopsis string `json:"synopsis" binding:"required"`
}

type UpdateBookCommand struct {
	bookRepository repository.BookRepository
}

func NewUpdateBookCommand(bookRepository repository.BookRepository) *UpdateBookCommand {
	return &UpdateBookCommand{bookRepository: bookRepository}
}

func (cmd *UpdateBookCommand) Execute(params UpdateBookParams) (*models.Book, error) {
	var finalErr error

	defer func() {
		if finalErr != nil {
			log.Printf("UpdateBookCommand failed, error: %v", finalErr)
		} else {
			log.Printf("UpdateBookCommand finished successfully.")
		}
	}()

	log.Printf("UpdateBookCommand initiated.")

	book, err := cmd.getBook(params.ID)

	if err != nil {
		finalErr = err
		return nil, err
	}

	updatedBook := cmd.updateBook(book, params)

	if err := cmd.bookRepository.Save(updatedBook.GetEntity()); err != nil {
		finalErr = err
		return nil, err
	}

	return updatedBook, nil
}

func (*UpdateBookCommand) updateBook(book *models.Book, params UpdateBookParams) *models.Book {
	book.Title = params.Title
	book.Category = params.Category
	book.Author = params.Author
	book.Synopsis = params.Synopsis

	return book
}

func (cmd *UpdateBookCommand) getBook(id int64) (*models.Book, error) {
	book, err := cmd.bookRepository.FindByID(id)

	if err != nil {
		return nil, err
	}

	return book, nil
}
