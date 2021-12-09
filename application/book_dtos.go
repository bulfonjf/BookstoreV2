package application

type BookDTO struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type CreateBookDTO struct {
	Title string `json:"title" validate:"required"`
}

type UpdateBookDTO struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
