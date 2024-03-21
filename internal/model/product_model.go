package model

type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

type GetProductsRequest struct {
	Page  int
	Limit int
	Name  string
}

type UpdateProductRequest struct {
	Name  string
	Stock *int
}

type CreateProductRequest struct {
	Name  string `json:"name_product" validate:"required"`
	Stock int    `json:"stock" validate:"required"`
}
