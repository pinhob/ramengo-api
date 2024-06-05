package protein_test

import (
	"testing"

	"github.com/pinhob/ramengo-api/protein"
)

var service protein.Service

func TestGetAll(t *testing.T) {
	ConfigureServiceHelper(t)
	proteins, err := service.GetAll()
	if err != nil {
		t.Errorf("got error calling GetAll function, %v", err)
	}
	mockDBSize := len(protein.MockDBProtein)

	if len(proteins) != mockDBSize {
		t.Errorf("got %v proteins, want %v", len(proteins), mockDBSize)
	}
}

func ConfigureServiceHelper(t *testing.T) {
	t.Helper()
	service = protein.Service{
		Repository: &protein.RepositoryData{
			MockData: protein.MockDBProtein,
		},
	}
}
