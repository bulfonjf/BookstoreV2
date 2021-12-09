package application

import "bookstore/domain"

func mapToInventoryDetailDTO(inventoryDetail domain.InventoryDetail) InventoryDetailDTO {
	return InventoryDetailDTO{
		BookDTO: mapToBookDTO(inventoryDetail.Book),
		Quantity: inventoryDetail.Quantity,
	}
}
