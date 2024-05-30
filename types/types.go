package types

type OrderRequest struct {
	BrothId   string
	ProteinId string
}

type OrderRespone struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
