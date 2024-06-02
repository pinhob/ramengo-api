package broth

type Broth struct {
	ID            int    `json:"id"`
	ImageInactive string `json:"imageInactive"`
	ImageActive   string `json:"imageActive"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Price         int    `json:"price"`
}
