package application

import (
	"bookstore/domain"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var (
	ErrCreateBook    = Error{Code: EINTERNAL, Message: "the book can't be created in the repository"}
	ErrNotFound      = Error{Code: ENOTFOUND, Message: "the book doesn't exist in the repository"}
	ErrInvalidBookID = Error{Code: EINVALID, Message: "Book id must be a valid uuid"}
)

func NewBookService(repository BookRepository) *BookService {
	return &BookService{
		repository: repository,
	}
}

type BookService struct {
	repository BookRepository
}

func (bs *BookService) CreateBook(createBookDTO CreateBookDTO) (BookDTO, error) {
	var book domain.Book
	book, err := bs.repository.GetBookByTitle(createBookDTO.Title)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return BookDTO{}, fmt.Errorf("creating book. Error: %w", err)
	} else if err != nil && errors.Is(err, ErrNotFound) {
		book, err = domain.NewBook(createBookDTO.Title)
		if err != nil && errors.Is(err, domain.ErrInvalidTitle) {
			return BookDTO{}, Error{Code: EINVALID, Message: err.Error()}
		}
		if cbError := bs.repository.CreateBook(book); cbError != nil {
			return BookDTO{}, cbError
		}
	}

	return mapToBookDTO(book), nil
}

func (bs *BookService) GetBooks() ([]BookDTO, error) {
	var booksDTO []BookDTO
	var books []domain.Book
	books, err := bs.repository.GetBooks()
	if err != nil {
		return booksDTO, err
	}

	for _, b := range books {
		booksDTO = append(booksDTO, mapToBookDTO(b))
	}

	return booksDTO, nil
}

func (bs *BookService) GetBookByID(id string) (BookDTO, error) {
	bookID, err := uuid.Parse(id)
	if err != nil {
		return BookDTO{}, ErrInvalidBookID
	}

	book, err := bs.repository.GetBookByID(bookID)
	if err != nil && errors.Is(err, ErrNotFound) {
		return BookDTO{}, err
	} else if err != nil {
		return BookDTO{}, fmt.Errorf("getting book by id from repository: Error: %w", err)
	}

	return mapToBookDTO(book), nil
}

func (bs *BookService) UpdateBook(updateBookDTO UpdateBookDTO) (BookDTO, error) {
	book, err := mapUpdateBookDTOToBook(updateBookDTO)
	if err != nil {
		return BookDTO{}, err
	}

	err = bs.repository.UpdateBook(book)
	if err != nil && errors.Is(err, ErrNotFound) {
		return BookDTO{}, err
	} else if err != nil {
		return BookDTO{}, fmt.Errorf("book service: can't update book. Error: %w", err)
	}

	return mapToBookDTO(book), nil
}

func (bs *BookService) DeleteBook(bookID string) error {
	id, err := ParseBookID(bookID)
	if err != nil {
		return err
	}

	err = bs.repository.DeleteBook(id)
	if err != nil && errors.Is(err, ErrNotFound) {
		return err
	} else if err != nil {
		return fmt.Errorf("book service: can't delete book. Error: %w", err)
	}

	return nil
}

func ParseBookID(id string) (uuid.UUID, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return uuid.UUID{}, ErrInvalidBookID
	}

	return parsedID, nil
}
