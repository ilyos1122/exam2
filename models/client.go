package models

type ClientPrimaryKey struct {
	Id string `json:"id"`
}

type CreateClient struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Phone       string `json:"phone"`
	Photo       string `json:"photo"`
	DateOfBirth string `json:"date_of_birth"`
}

type Client struct {
	Id          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Phone       string `json:"phone"`
	Photo       string `json:"photo"`
	DateOfBirth string `json:"date_of_birth"`
	CreatedAT   string `json:"created_at"`
}

type UpdateClient struct {
	Id    string `json:"id"`
	Phone string `json:"phone"`
	Photo string `json:"photo"`
}

type GetListClientRequest struct {
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
	Search string `json:"search"`
}

type GetListClientResponse struct {
	Count   int64     `json:"count"`
	Clients []*Client `json:"clients"`
}
