package model

// Response is standard api response model.
type Response struct {
	Status     int         `json:"status"`
	Message    string      `json:"message"`
	Data       any         `json:"data"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// Pagination is pagination response model.
type Pagination struct {
	Total       int `json:"total"`
	Limit       int `json:"limit"`
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
}
