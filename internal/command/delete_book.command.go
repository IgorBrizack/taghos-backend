package command

import (
	"log"

	"github.com/IgorBrizack/taghos-backend/internal/repository"
)

type DeleteBookParams struct {
	ID int64
}

type DeleteBookCommand struct {
	bookRepository repository.BookRepository
}

func NewDeleteBookCommand(bookRepository repository.BookRepository) *DeleteBookCommand {
	return &DeleteBookCommand{bookRepository: bookRepository}
}

func (cmd *DeleteBookCommand) Execute(params DeleteBookParams) error {
	var finalErr error

	defer func() {
		if finalErr != nil {
			log.Printf("DeleteBookCommand failed, error: %v", finalErr)
		} else {
			log.Printf("DeleteBookCommand finished successfully.")
		}
	}()

	log.Printf("DeleteBookCommand initiated for book ID: %d", params.ID)

	err := cmd.bookRepository.Delete(params.ID)
	if err != nil {
		finalErr = err
		return err
	}

	return nil
}
