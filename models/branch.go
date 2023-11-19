package models

type BranchPrimaryKey struct {
	Id string `json:"id"`
}

type CreateBranch struct {
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Photo     string `json:"photo"`
	WorkStart string `json:"work_start_hour"`
	WorkEnd   string `json:"work_end_hour"`
	Address   string `json:"address"`
}

type Branch struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	Photo         string `json:"photo"`
	WorkStart     string `json:"work_start_hour"`
	WorkEnd       string `json:"work_end_hour"`
	Address       string `json:"address"`
	DeliveryPrice int    `json:"delivery_price"`
	Status        string `json:"status"`
}

type UpdateBranch struct {
	Id            string `json:"id"`
	WorkStart     string `json:"work_start_hour"`
	WorkEnd       string `json:"work_end_hour"`
	DeliveryPrice int    `json:"delivery_price"`
	Phone         string `json:"phone"`
	Photo         string `json:"photo"`
}

type GetListBranchRequest struct {
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
	Search string `json:"search"`
	Query  string `json:"query"`
}

type GetListBranchResponse struct {
	Count    int       `json:"count"`
	Branches []*Branch `json:"branches"`
}
