package inmem

type InMemRepository struct {
	books     []inmemBook
	inventory map[string]uint
}

func NewInMemoryRepository(dns string) *InMemRepository {
	return &InMemRepository{
		books:     make([]inmemBook, 0),
		inventory: make(map[string]uint),
	}
}

func (i *InMemRepository) Open() error {
	return nil
}

func (i *InMemRepository) Close() error {
	return nil
}
