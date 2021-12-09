package application

import "bookstore/domain"

func mapToBookDTO(book domain.Book) BookDTO {
	return BookDTO{
		ID:    book.ID.String(),
		Title: book.Title,
	}
}

func mapUpdateBookDTOToBook(book UpdateBookDTO) (domain.Book, error) {
	id, err := ParseBookID(book.ID)
	if err != nil {
		return domain.Book{}, err
	}

	return domain.Book{
		ID:    id,
		Title: book.Title,
	}, nil
}

func mapToBook(bookDTO BookDTO) (domain.Book, error) {
	id, err := ParseBookID(bookDTO.ID)
	if err != nil {
		return domain.Book{}, err
	}

	return domain.Book{
		ID:    id,
		Title: bookDTO.Title,
	}, nil
}
