package dish

import "errors"

type Reader interface {
	GetDish(brothId, proteinId int) (Dish, error)
}

type Repository interface {
	Reader
}

type RepositoryData struct {
	MockData []Dish
}

func (r *RepositoryData) GetDish(brothId, proteinId int) (Dish, error) {
	for _, d := range r.MockData {
		if d.BrothId == brothId && d.ProteinId == proteinId {
			return d, nil
		}
	}

	return Dish{}, errors.New("we don't have any dish with the given broth and protein ids")
}
