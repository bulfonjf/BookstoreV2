package inmem

import (
	"bookstore/application"
	"bookstore/domain"

	"github.com/google/uuid"
)

type inmemBook struct {
	id    string
	title string
}

func (i *InMemRepository) CreateBook(book domain.Book) error {
	newBook := inmemBook{
		id:    book.ID.String(),
		title: book.Title,
	}

	i.books = append(i.books, newBook)

	return nil
}

func (i *InMemRepository) GetBooks() ([]domain.Book, error) {
	var books []domain.Book

	for _, b := range i.books {
		parsedID, err := application.ParseBookID(b.id)
		if err != nil {
			return []domain.Book{}, err
		}

		books = append(books, domain.Book{ID: parsedID, Title: b.title})
	}

	return books, nil
}

func (i *InMemRepository) GetBookByID(id uuid.UUID) (domain.Book, error) {
	bookFound := domain.Book{}
	bookIndex := i.getBookIndex(id)
	if bookIndex < 0 {
		return domain.Book{}, application.ErrNotFound
	}

	b := i.books[bookIndex]
	parsedID, err := application.ParseBookID(b.id)
	if err != nil {
		return domain.Book{}, err
	}

	bookFound = domain.Book{ID: parsedID, Title: b.title}

	return bookFound, nil
}

func (i *InMemRepository) GetBookByTitle(title string) (domain.Book, error) {
	for _, b := range i.books {
		if title == b.title {
			parsedID, err := application.ParseBookID(b.id)
			if err != nil {
				return domain.Book{}, err
			}

			return domain.Book{
				ID: parsedID,
				Title: b.title,
			}, nil
		}
	}

	return domain.Book{}, application.ErrNotFound
}

func (i *InMemRepository) UpdateBook(book domain.Book) error {
	bookIndex := i.getBookIndex(book.ID)
	if bookIndex < 0 {
		return application.ErrNotFound
	} else {
		i.books[bookIndex] = inmemBook{id: book.ID.String(), title: book.Title}
		return nil
	}
}

func (i *InMemRepository) DeleteBook(id uuid.UUID) error {
	_, err := i.GetBookByID(id)
	if err != nil {
		return err
	}

	bookIndex := i.getBookIndex(id)
	i.books = deleteBook(i.books, bookIndex)

	return nil
}

func (i *InMemRepository) getBookIndex(id uuid.UUID) int {
	for index, b := range i.books {
		if id.String() == b.id {
			return index
		}

	}

	return -1
}

func deleteBook(books []inmemBook, indexToRemove int) []inmemBook {
	currentLength := len(books)

	if currentLength == 0 {
		return books
	}

	lastItem := currentLength - 1

	switch true {
	case indexToRemove == lastItem:
		books = books[:lastItem]
	case indexToRemove > 0 && indexToRemove < lastItem:
		books[indexToRemove] = books[lastItem]
		books = books[:lastItem]
	case indexToRemove == 0:
		books = books[1:]
	}

	return books
}
