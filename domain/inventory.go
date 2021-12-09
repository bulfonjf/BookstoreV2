package domain

type InventoryDetail struct {
	Book
	Quantity uint
}

func NewInventoryDetail(book Book, q uint) InventoryDetail {
	return InventoryDetail{
		Book: book,
		Quantity: q,
	}
}