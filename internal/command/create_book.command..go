package command

import (
	"log"

	"github.com/IgorBrizack/taghos-backend/internal/models"
	"github.com/IgorBrizack/taghos-backend/internal/repository"
)

type CreateBookParams struct {
	Title    string `json:"title" binding:"required"`
	Category string `json:"category" binding:"required"`
	Author   string `json:"author" binding:"required"`
	Synopsis string `json:"synopsis" binding:"required"`
}

type CreateBookCommand struct {
	bookRepository repository.BookRepository
}

func NewCreateBookCommand(bookRepository repository.BookRepository) *CreateBookCommand {
	return &CreateBookCommand{bookRepository: bookRepository}
}

func (cmd *CreateBookCommand) Execute(params CreateBookParams) (*models.Book, error) {
	var finalErr error

	defer func() {
		if finalErr != nil {
			log.Printf("CreateBookCommand failed, error: %v", finalErr)
		} else {
			log.Printf("CreateBookCommand finished successfully.")
		}
	}()

	log.Printf("CreateBookCommandt command initiated")

	book := models.Book{
		Title:    params.Title,
		Category: params.Category,
		Author:   params.Author,
		Synopsis: params.Synopsis,
	}

	if err := cmd.bookRepository.Create(&book); err != nil {
		finalErr = err
		return nil, err
	}

	return &book, nil
}
