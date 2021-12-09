package application

type InventoryDetailDTO struct {
	BookDTO
	Quantity uint
}

type InventoryDTO []InventoryDetailDTO