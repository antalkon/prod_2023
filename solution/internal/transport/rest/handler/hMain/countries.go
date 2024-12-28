package hmain

import (
	"net/http"
	"sort"

	pgmain "github.com/antalkon/prod_2023/internal/db/pgMain"
	"github.com/antalkon/prod_2023/internal/models"
	"github.com/labstack/echo/v4"
)

func Countries(c echo.Context) error {
	region := c.QueryParam("region")

	var (
		countries []models.Countries
		err       error
	)

	// Получение данных
	if region == "" {
		countries, err = pgmain.GetAllCountries()
	} else {
		countries, err = pgmain.GetCountriesByRegion(region)
	}

	// Обработка ошибок
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"reason": "Ошибка обработки запроса: " + err.Error(),
		})
	}

	// Сортировка стран по двухбуквенному коду (alpha2)
	sort.Slice(countries, func(i, j int) bool {
		return countries[i].Alpha2 < countries[j].Alpha2
	})

	// Возвращаем успешный результат
	return c.JSON(http.StatusOK, countries)
}
