package model

type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

type GetProductListRequest struct {
	Page  int
	Limit int
	Name  string
}

type UpdateProductRequest struct {
	Name  string `json:"name"`
	Stock *int   `json:"stock"`
}

type CreateProductRequest struct {
	Name  string `json:"name" validate:"required"`
	Stock int    `json:"stock" validate:"required"`
}
