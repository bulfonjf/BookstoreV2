package application

import (
	"bookstore/domain"

	"github.com/google/uuid"
)

type BookRepository interface {
	CreateBook(book domain.Book) error
	DeleteBook(id uuid.UUID) error
	GetBooks() ([]domain.Book, error)
	GetBookByID(id uuid.UUID) (domain.Book, error)
	GetBookByTitle(title string) (domain.Book, error) 
	UpdateBook(book domain.Book) error
}
