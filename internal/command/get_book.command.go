package command

import (
	"log"

	"github.com/IgorBrizack/taghos-backend/internal/models"
	"github.com/IgorBrizack/taghos-backend/internal/repository"
)

type GetBookParams struct {
	ID int64
}

type GetBookCommand struct {
	bookRepository repository.BookRepository
}

func NewGetBookCommand(bookRepository repository.BookRepository) *GetBookCommand {
	return &GetBookCommand{bookRepository: bookRepository}
}

func (cmd *GetBookCommand) Execute(params GetBookParams) (*models.Book, error) {
	var finalErr error

	defer func() {
		if finalErr != nil {
			log.Printf("GetBookCommand failed, error: %v", finalErr)
		} else {
			log.Printf("GetBookCommand finished successfully.")
		}
	}()

	log.Printf("GetBookCommand initiated for book ID: %d", params.ID)

	book, err := cmd.bookRepository.FindByID(params.ID)

	if err != nil {
		finalErr = err
		return nil, err
	}

	return book, nil
}
