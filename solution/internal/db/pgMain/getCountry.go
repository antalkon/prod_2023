package pgmain

import (
	"fmt"

	"github.com/antalkon/prod_2023/internal/models"
	"github.com/antalkon/prod_2023/pkg/pggorm"
)

func GetCountryByAlpha2(alpha2 string) (models.Countries, error) {
	db := pggorm.DB
	if db == nil {
		return models.Countries{}, fmt.Errorf("database connection is not initialized")
	}

	var country models.Countries
	if err := db.Where("alpha2 =?", alpha2).First(&country).Error; err != nil {
		return models.Countries{}, fmt.Errorf("failed to get country by alpha2: %w", err)
	}

	return country, nil
}
