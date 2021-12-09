package application

import "bookstore/domain"

type InventoryRepository interface {
	AddBook(book domain.Book, q uint) error
	GetInventory() (map[string]uint, error)
}