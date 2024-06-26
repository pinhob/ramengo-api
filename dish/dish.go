package dish

type Dish struct {
	ID        int
	Name      string
	Image     string
	BrothId   int
	ProteinId int
}

var MockDBDishes = []Dish{
	{ID: 7, Name: "Salt and Chasu Ramen", Image: "https://tech.redventures.com.br/icons/ramen/ramenChasu.png", BrothId: 1, ProteinId: 1},
	{ID: 8, Name: "Salt and Yasai Vegetable Ramen", Image: "https://tech.redventures.com.br/icons/ramen/ramenYasai%20Vegetarian.png", BrothId: 1, ProteinId: 2},
	{ID: 9, Name: "Salt and Karaague Ramen", Image: "https://tech.redventures.com.br/icons/ramen/ramenKaraague.png", BrothId: 1, ProteinId: 3},
	{ID: 1, Name: "Shoyu and Chasu Ramen", Image: "https://tech.redventures.com.br/icons/ramen/ramenChasu.png", BrothId: 2, ProteinId: 1},
	{ID: 2, Name: "Shoyu and Yasai Vegetable Ramen", Image: "https://tech.redventures.com.br/icons/ramen/ramenYasai%20Vegetarian.png", BrothId: 2, ProteinId: 2},
	{ID: 3, Name: "Shoyu and Karague Ramen", Image: "https://tech.redventures.com.br/icons/ramen/ramenKaraague.png", BrothId: 2, ProteinId: 3},
	{ID: 4, Name: "Miso and Chasu Ramen", Image: "https://tech.redventures.com.br/icons/ramen/ramenChasu.png", BrothId: 3, ProteinId: 1},
	{ID: 5, Name: "Miso and Yasai Vegetable Ramen", Image: "https://tech.redventures.com.br/icons/ramen/ramenYasai%20Vegetarian.png", BrothId: 3, ProteinId: 2},
	{ID: 6, Name: "Miso and Kaarague Ramen", Image: "https://tech.redventures.com.br/icons/ramen/ramenKaraague.png", BrothId: 3, ProteinId: 3},
}
