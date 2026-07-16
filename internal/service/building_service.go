package service

import (
	"os"

	"f4b1.dev/clicker-backend/internal/models"
	"github.com/gocarina/gocsv"
)

type BuildingService struct {
	Buildings map[string]*models.Building
}

func NewBuildingService() (*BuildingService, error) {
	file, err := os.Open("data/buildings.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var buildings []*models.Building
	err = gocsv.UnmarshalFile(file, &buildings)
	if err != nil {
		return nil, err
	}

	buildingsMap := make(map[string]*models.Building)
	for _, b := range buildings {
		buildingsMap[b.Id] = b
	}

	return &BuildingService{Buildings: buildingsMap}, nil
}
