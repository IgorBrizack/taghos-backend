package repository

import (
	"github.com/IgorBrizack/taghos-backend/internal/models"
	"gorm.io/gorm"
)

type BookRepository interface {
	Create(book *models.Book) error
	Save(book *models.Book) error
	Delete(id int64) error
	FindByID(id int64) (*models.Book, error)
	FindAll() ([]*models.Book, error)
}

type BookRepositoryImpl struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &BookRepositoryImpl{DB: db}
}

func (r *BookRepositoryImpl) Create(book *models.Book) error {
	return r.DB.Create(book).Error
}

func (r *BookRepositoryImpl) Save(book *models.Book) error {
	return r.DB.Save(book).Error
}

func (r *BookRepositoryImpl) Delete(id int64) error {
	return r.DB.Delete(&models.Book{}, id).Error
}

func (r *BookRepositoryImpl) FindByID(id int64) (*models.Book, error) {
	var book models.Book
	if err := r.DB.First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepositoryImpl) FindAll() ([]*models.Book, error) {
	var books []*models.Book
	if err := r.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
