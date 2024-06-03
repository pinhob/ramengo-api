package broth

type Reader interface {
	GetAll() ([]Broth, error)
}
type Repository interface {
	Reader
}

type RepositoryData struct {
	MockData []Broth
}

func (r *RepositoryData) GetAll() ([]Broth, error) {
	return r.MockData, nil
}
