package command

import (
	"log"

	"github.com/IgorBrizack/taghos-backend/internal/models"
	"github.com/IgorBrizack/taghos-backend/internal/repository"
)

type GetAllBooksCommand struct {
	bookRepository repository.BookRepository
}

func NewGetAllBooksCommand(bookRepository repository.BookRepository) *GetAllBooksCommand {
	return &GetAllBooksCommand{bookRepository: bookRepository}
}

func (cmd *GetAllBooksCommand) Execute() ([]*models.Book, error) {
	var finalErr error

	defer func() {
		if finalErr != nil {
			log.Printf("GetAllBooksCommand failed, error: %v", finalErr)
		} else {
			log.Printf("GetAllBooksCommand finished successfully.")
		}
	}()

	log.Printf("GetAllBooksCommand initiated.")

	books, err := cmd.bookRepository.FindAll()

	if err != nil {
		finalErr = err
		return nil, err
	}

	return books, nil
}
