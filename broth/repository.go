package broth

type Reader interface {
}
type Repository interface {
	GetAll() ([]Broth, error)
}

func GetAll() ([]Broth, error) {
	return MockDBTable, nil
}

var MockDBTable = []Broth{
	{ID: 1, ImageInactive: "https://tech.redventures.com.br/icons/salt/inactive.svg", ImageActive: "https://tech.redventures.com.br/icons/salt/active.svg", Name: "Salt", Description: "Simple like the seawater, nothing more", Price: 10},
	{ID: 2, ImageInactive: "https://tech.redventures.com.br/icons/shoyu/inactive.svg", ImageActive: "https://tech.redventures.com.br/icons/shoyu/active.svg", Name: "Shoyu", Description: "The good old and traditional soy sauce", Price: 10},
	{ID: 3, ImageInactive: "https://tech.redventures.com.br/icons/miso/inactive.svg", ImageActive: "https://tech.redventures.com.br/icons/miso/active.svg", Name: "Miso", Description: "Paste made of fermented soybeans", Price: 12},
}
