package models

type ProductPrimaryKey struct {
	Id string `json:"id"`
}

type CreateProduct struct {
	Id          string  `json:"id"`
	ProductID   string  `json:"product_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"product_image"`
	CategoryId  string  `json:"category_id"`
}

type Product struct {
	Id          string      `json:"id"`
	ProductID   string      `json:"product_id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Price       float64     `json:"price"`
	Image       string      `json:"product_image"`
	CategoryId  string      `json:"category_id"`
	Category    Category `json:"category"`
	CreatedAt   string      `json:"created_at"`
	UpdatedAt   string      `json:"updated_at"`
}

type UpdateProduct struct {
	Id          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"product_image"`
	CategoryId  string  `json:"category_id"`
}

type GetListProductRequest struct {
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
	Search string `json:"search"`
}

type GetListProductResponse struct {
	Count    int        `json:"count"`
	Products []*Product `json:"products"`
}
