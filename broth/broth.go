package broth

type Broth struct {
	ID            int
	ImageInactive string
	ImageActive   string
	Name          string
	Description   string
	Price         int
}

var BrothsTable = []Broth{
	{ID: 1, ImageInactive: "#", ImageActive: "#", Name: "Salt", Description: "Simple like the seawater, nothing more", Price: 10},
	{ID: 2, ImageInactive: "#", ImageActive: "#", Name: "Shoyu", Description: "The good old and traditional soy sauce", Price: 10},
	{ID: 3, ImageInactive: "#", ImageActive: "#", Name: "Miso", Description: "Paste made of fermented soybeans", Price: 12},
}
