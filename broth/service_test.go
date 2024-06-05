package broth_test

import (
	"testing"

	"github.com/pinhob/ramengo-api/broth"
)

var service broth.Service

func TestGetAll(t *testing.T) {
	ConfigureServiceHelper(t)
	broths, err := service.GetAll()
	if err != nil {
		t.Errorf("got error calling GetAll function, %v", err)
	}
	mockDBSize := len(broth.MockDBTable)

	if len(broths) != mockDBSize {
		t.Errorf("got %v broths, want %v", len(broths), mockDBSize)
	}
}

func ConfigureServiceHelper(t *testing.T) {
	t.Helper()
	service = broth.Service{
		Repository: &broth.RepositoryData{
			MockData: broth.MockDBTable,
		},
	}
}
