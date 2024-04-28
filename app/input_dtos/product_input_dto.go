package input_dtos

type CreateProductInputDTO struct {
	Name        string    `json:"name"        binding:"required"`
	Description string    `json:"description" binding:"required"`
	Price       float64   `json:"price"       binding:"required"`
	Quantity    int       `json:"quantity"    binding:"required"`
	Active      bool      `json:"active"`
}

type UpdateProductInputDTO struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	Active      bool      `json:"active"`
}
