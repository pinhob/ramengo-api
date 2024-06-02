package broth

import "errors"

type Service struct {
	Repository Repository
}

func (s Service) GetAll() ([]Broth, error) {
	broths, err := GetAll()
	if err != nil {
		return nil, errors.New("cannot get broths from repository")
	}

	return broths, nil
}
