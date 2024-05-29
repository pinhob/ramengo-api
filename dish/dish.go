package dish

import "errors"

type Dish struct {
	ID        int
	Name      string
	Image     string
	BrothId   int
	ProteinId int
}

var DishTable = []Dish{
	{ID: 1, Name: "Shoyu and Chasu Ramen", Image: "#", BrothId: 1, ProteinId: 1},
	{ID: 2, Name: "Shoyu and Yasai Vegetable Ramen", Image: "#", BrothId: 1, ProteinId: 2},
	{ID: 3, Name: "Shoyu and Karague Ramen", Image: "#", BrothId: 1, ProteinId: 3},
	{ID: 4, Name: "Miso and Chasu Ramen", Image: "#", BrothId: 3, ProteinId: 1},
	{ID: 5, Name: "Miso and Yasai Vegetable Ramen", Image: "#", BrothId: 3, ProteinId: 2},
	{ID: 6, Name: "Miso and Kaarague Ramen", Image: "#", BrothId: 3, ProteinId: 3},
	{ID: 7, Name: "Salt and Chasu Ramen", Image: "#", BrothId: 1, ProteinId: 1},
	{ID: 8, Name: "Salt and Yasai Vegetable Ramen", Image: "#", BrothId: 1, ProteinId: 2},
	{ID: 9, Name: "Salt and Karaague Ramen", Image: "#", BrothId: 1, ProteinId: 3},
}

func GetDish(brothId, proteinId int) (error, Dish) {
	for _, d := range DishTable {
		if d.BrothId == brothId && d.ProteinId == proteinId {
			return nil, d
		}
	}

	return errors.New("we don't have any dish with the given broth and protein ids"), Dish{}
}
