package command

import (
	"log"

	"github.com/IgorBrizack/taghos-backend/internal/models"
	"github.com/IgorBrizack/taghos-backend/internal/repository"
)

type CreateBookParams struct {
	Title    string
	Category string
	Author   string
	Synopsis string
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
