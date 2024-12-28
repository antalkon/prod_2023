package pgmain

import (
	"fmt"

	"github.com/antalkon/prod_2023/internal/models"
	"github.com/antalkon/prod_2023/pkg/pggorm"
)

func GetCountriesByRegion(region string) ([]models.Countries, error) {
	db := pggorm.DB
	if db == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	var countries []models.Countries
	err := db.Where("region =?", region).Find(&countries).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get countries by region: %w", err)
	}
	return countries, nil

}

func GetAllCountries() ([]models.Countries, error) {
	db := pggorm.DB
	if db == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	var countries []models.Countries
	err := db.Find(&countries).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get all countries: %w", err)
	}
	return countries, nil
}
