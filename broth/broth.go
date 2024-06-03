package broth

type Broth struct {
	ID            int    `json:"id"`
	ImageInactive string `json:"imageInactive"`
	ImageActive   string `json:"imageActive"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Price         int    `json:"price"`
}

var MockDBTable = []Broth{
	{ID: 1, ImageInactive: "https://tech.redventures.com.br/icons/salt/inactive.svg", ImageActive: "https://tech.redventures.com.br/icons/salt/active.svg", Name: "Salt", Description: "Simple like the seawater, nothing more", Price: 10},
	{ID: 2, ImageInactive: "https://tech.redventures.com.br/icons/shoyu/inactive.svg", ImageActive: "https://tech.redventures.com.br/icons/shoyu/active.svg", Name: "Shoyu", Description: "The good old and traditional soy sauce", Price: 10},
	{ID: 3, ImageInactive: "https://tech.redventures.com.br/icons/miso/inactive.svg", ImageActive: "https://tech.redventures.com.br/icons/miso/active.svg", Name: "Miso", Description: "Paste made of fermented soybeans", Price: 12},
}
