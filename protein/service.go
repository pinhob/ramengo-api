package protein

import "errors"

type Service struct {
	Repository Repository
}

func (s Service) GetAll() ([]Protein, error) {
	broths, err := s.Repository.GetAll()
	if err != nil {
		return nil, errors.New("cannot get broths from repository")
	}

	return broths, nil
}
