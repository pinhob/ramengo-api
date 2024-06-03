package dish

import "errors"

type Service struct {
	Repository Repository
}

func (s Service) GetDishByCondiments(brothId, proteinId int) (Dish, error) {
	dish, err := s.Repository.GetDish(brothId, proteinId)
	if err != nil {
		return Dish{}, errors.New("cannot get dish to order")
	}

	return dish, nil
}
