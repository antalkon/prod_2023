package hmain

import (
	"net/http"

	pgmain "github.com/antalkon/prod_2023/internal/db/pgMain"
	"github.com/labstack/echo/v4"
)

func GetCountry(c echo.Context) error {
	alpha2 := c.Param("alpha2")
	if alpha2 == "" {
		return c.JSON(404, map[string]string{
			"reason": "Не указан alpha2-код страны",
		})
	}
	country, err := pgmain.GetCountryByAlpha2(alpha2)
	if err != nil {
		return c.JSON(404, map[string]string{
			"reason": "Ошибка обработки запроса: " + err.Error(),
		})
	}
	return c.JSON(http.StatusOK, country)

}
