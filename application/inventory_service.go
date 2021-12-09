package application

import (
	"bookstore/domain"
	"errors"
	"fmt"
)

func NewInventoryService(repository InventoryRepository, bookRepository BookRepository) *InventoryService {
	return &InventoryService{
		repository:     repository,
		bookRepository: bookRepository,
	}
}

type InventoryService struct {
	repository     InventoryRepository
	bookRepository BookRepository
}

func (i *InventoryService) AddBook(bookDTO BookDTO) error {
	book, err := mapToBook(bookDTO)
	if err != nil {
		return fmt.Errorf("adding book to inventory. Error: %w", err)
	}

	if err = i.repository.AddBook(book, 1); err != nil {
		return fmt.Errorf("adding book to inventory. Error: %w", err)
	}

	return nil
}

func (i *InventoryService) GetInventory() ([]InventoryDetailDTO, error) {
	inventoryDTO := make([]InventoryDetailDTO, 0)

	inventory, err := i.repository.GetInventory()
	if err != nil {
		return []InventoryDetailDTO{}, fmt.Errorf("getting the inventory. Error: %w", err)
	}

	for k, v := range inventory {
		bookID, err := ParseBookID(k)
		if err != nil {
			return []InventoryDetailDTO{}, fmt.Errorf("getting the inventory and parsing book ID. Error: %w", err)
		}

		book, err := i.bookRepository.GetBookByID(bookID)
		if err != nil && errors.Is(err, ErrNotFound) {
			// todo delete book from inventory
			continue
		} else if err != nil {
			return []InventoryDetailDTO{}, fmt.Errorf("getting the inventory and getting books details. Error: %w", err)
		}

		inventoryDetail := domain.NewInventoryDetail(book, v)

		inventoryDTO = append(inventoryDTO, mapToInventoryDetailDTO(inventoryDetail))
	}

	return inventoryDTO, nil
}
