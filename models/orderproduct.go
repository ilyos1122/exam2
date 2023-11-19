package models

type OrderProductPrimaryKey struct {
	OrderProductId string `json:"order_product_id"`
}

type OrderProduct struct {
	OrderProductId string `json:"order_product_id"`
	OrderID        string `json:"order_id"`
	ProductID      string `json:"product_id"`
	DiscountType   string `json:"discount_type"`
	DiscountAmount int    `json:"discount_amount"`
	Quantity       int    `json:"quantity"`
	Price          int    `json:"price"`
	Sum            int    `json:"sum"`
	CreatedAt      string `json:"created_at"`
}

type CreateOrderProduct struct {
	OrderID        string `json:"order_id"`
	ProductID      string `json:"product_id"`
	DiscountType   string `json:"discount_type"`
	DiscountAmount int    `json:"discount_amount"`
	Quantity       int    `json:"quantity"`
}
