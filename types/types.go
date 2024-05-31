package types

type OrderRequest struct {
	BrothId   string
	ProteinId string
}

type OrderRespone struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
