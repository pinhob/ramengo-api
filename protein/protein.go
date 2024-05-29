package protein

type Protein struct {
	ID            int
	ImageInactive string
	ImageActive   string
	Name          string
	Description   string
	Price         int
}

var ProteinsTable = []Protein{
	{ID: 1, ImageInactive: "#", ImageActive: "#", Name: "Chasu", Description: "A sliced flavourful pork meat with a selection of season vegetables.", Price: 10},
	{ID: 2, ImageInactive: "#", ImageActive: "#", Name: "Yasai Vegetarian", Description: "A delicious vegetarian lamen with a selection of season vegetables.", Price: 10},
	{ID: 3, ImageInactive: "#", ImageActive: "#", Name: "Karaague", Description: "Three units of fried chicken, moyashi, ajitama egg and other vegetables.", Price: 12},
}
