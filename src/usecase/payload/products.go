package payload

type GetProductsRequest struct {
	Find   string `json:"find"`
	Sort   string `json:"sort"`
	Filter string `json:"filter"`
}

type CreateProductRequest struct {
	Name  string  `json:"name" validate:"required"`
	Types string  `json:"types" validate:"required"`
	Price float64 `json:"price" validate:"required"`
}
