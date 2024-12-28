package hme

import (
	"net/http"
	"strconv"

	pgdata "github.com/antalkon/prod_2023/internal/db/pgData"
	"github.com/labstack/echo/v4"
)

func MyProfile(c echo.Context) error {
	// Получение user_id из контекста
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"reason": "Invalid or missing user ID",
		})
	}

	// Получение пользователя из базы данных
	user, err := pgdata.GetUserById(strconv.Itoa(userID))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"reason": "Invalid token or user not found",
		})
	}

	// Формирование ответа
	profile := map[string]interface{}{
		"login":       user.Login,
		"email":       user.Email,
		"countryCode": user.CountryCode,
		"isPublic":    user.IsPublic,
		"phone":       user.Phone,
		"image":       user.Image,
	}

	return c.JSON(http.StatusOK, profile)
}
