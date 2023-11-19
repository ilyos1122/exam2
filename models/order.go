package models


type OrderPrimaryKey struct {
	ID              string       `json:"id"`
}

type Order struct {
	ID              string       `json:"id"`
	OrderID         string       `json:"order_id"`
	ClientID        string       `json:"client_id"`
	BranchID        string       `json:"branch_id"`
	DeliveryAddress string       `json:"delivery_address"`
	DeliveryPrice   int          `json:"delivery_price"`
	TotalCount      int          `json:"total_count"`
	TotalPrice      int          `json:"total_price"`
	OrderStatus     string       `json:"status"`
	OrderProduct   	[]OrderProduct `json:"order_products"`
	CreatedAt       string       `json:"created_at"`
	UpdatedAt       string       `json:"updated_at"`
}

type CreateOrder struct {
	ClientID        string `json:"client_id"`
	BranchID        string `json:"branch_id"`
	DeliveryAddress string `json:"delivery_address"`
}

type UpdateOrder struct {
	DeliveryAddress string `json:"delivery_address"`
	BranchID        string `json:"branch_id"`
	OrderID         string `json:"order_id"`
}

type GetListOrderRequest struct {
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
	Search string `json:"search"`
}

type GetListOrderResponse struct {
	Count  int      `json:"count"`
	Orders []*Order `json:"products"`
}

type ChangeOrderStatus struct {
	ID              string       `json:"id"`
	Status string `json:"order_status"`
}
