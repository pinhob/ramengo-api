package protein

type Protein struct {
	ID            int    `json:"id"`
	ImageInactive string `json:"imageInactive"`
	ImageActive   string `json:"imageActive"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Price         int    `json:"price"`
}

var ProteinsTable = []Protein{
	{ID: 1, ImageInactive: "#", ImageActive: "#", Name: "Chasu", Description: "A sliced flavourful pork meat with a selection of season vegetables.", Price: 10},
	{ID: 2, ImageInactive: "#", ImageActive: "#", Name: "Yasai Vegetarian", Description: "A delicious vegetarian lamen with a selection of season vegetables.", Price: 10},
	{ID: 3, ImageInactive: "#", ImageActive: "#", Name: "Karaague", Description: "Three units of fried chicken, moyashi, ajitama egg and other vegetables.", Price: 12},
}
