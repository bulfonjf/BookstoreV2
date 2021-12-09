package inmem

import (
	"bookstore/domain"
)

type Inventory struct {
}

func (i *InMemRepository) AddBook(book domain.Book, q uint) error {
	i.inventory[book.ID.String()] += q

	return nil
}

func (i *InMemRepository) GetInventory() (map[string]uint, error) {
	return i.inventory, nil
}
