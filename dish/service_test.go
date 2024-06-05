package dish_test

import (
	"testing"

	"github.com/pinhob/ramengo-api/dish"
)

var service dish.Service

func TestGetDish(t *testing.T) {
	ConfigureServiceHelper(t)
	brothId := 2
	proteinId := 3
	want := "Shoyu and Karague Ramen"

	dish, err := service.GetDishByCondiments(brothId, proteinId)
	if err != nil {
		t.Errorf("got error calling GetDishByCondiments function, %v", err)
	}

	if dish.Name != want {
		t.Errorf("got %s dish, want %s", dish.Name, want)
	}
}

func ConfigureServiceHelper(t *testing.T) {
	t.Helper()
	service = dish.Service{
		Repository: &dish.RepositoryData{
			MockData: dish.MockDBDishes,
		},
	}
}
