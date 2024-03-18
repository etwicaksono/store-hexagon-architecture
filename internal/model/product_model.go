package model

type GetProductRequest struct {
	Page  int
	Limit int
	Name  string
}

type UpdateProductRequest struct {
	Name  string
	Stock *int
}
