package protein

type Reader interface {
	GetAll() ([]Protein, error)
}
type Repository interface {
	Reader
}

type RepositoryData struct {
	MockData []Protein
}

func (r *RepositoryData) GetAll() ([]Protein, error) {
	return r.MockData, nil
}
