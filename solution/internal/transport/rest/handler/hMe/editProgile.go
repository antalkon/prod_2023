package hme

import (
	"errors"
	"net/http"
	"strconv"

	pgdata "github.com/antalkon/prod_2023/internal/db/pgData"
	"github.com/labstack/echo/v4"
)

func EditMyProfile(c echo.Context) error {
	// Получение user_id из контекста
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"reason": "Unauthorized: token is missing or invalid",
		})
	}

	// Получение текущего пользователя из базы данных
	user, err := pgdata.GetUserById(strconv.Itoa(userID))
	if err != nil {
		if errors.Is(err, pgdata.ErrUserNotFound) {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"reason": "User not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"reason": "Failed to fetch user",
		})
	}

	// Привязка данных из тела запроса
	var updatedData struct {
		CountryCode string `json:"countryCode,omitempty"`
		IsPublic    *bool  `json:"isPublic,omitempty"`
		Phone       string `json:"phone,omitempty"`
		Image       string `json:"image,omitempty"`
	}
	if err := c.Bind(&updatedData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"reason": "Invalid request body",
		})
	}

	// Валидация и обновление переданных данных
	if updatedData.CountryCode != "" {
		if len(updatedData.CountryCode) != 2 {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"reason": "Invalid country code: must be exactly 2 characters",
			})
		}
		user.CountryCode = updatedData.CountryCode
	}

	if updatedData.IsPublic != nil {
		user.IsPublic = *updatedData.IsPublic
	}

	if updatedData.Phone != "" {
		if len(updatedData.Phone) > 15 {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"reason": "Phone number exceeds maximum length of 15 characters",
			})
		}
		user.Phone = updatedData.Phone
	}

	if updatedData.Image != "" {
		if len(updatedData.Image) > 255 {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"reason": "Image URL exceeds maximum length of 255 characters",
			})
		}
		user.Image = updatedData.Image
	}

	// Сохранение обновленного профиля
	err = pgdata.UpdateUserProfile(user)
	if err != nil {
		if errors.Is(err, pgdata.ErrDuplicateEntry) {
			return c.JSON(http.StatusConflict, map[string]string{
				"reason": "Phone number is already in use",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"reason": "Failed to update profile",
		})
	}

	// Формируем успешный ответ
	return c.JSON(http.StatusOK, map[string]interface{}{
		"login":       user.Login,
		"email":       user.Email,
		"countryCode": user.CountryCode,
		"isPublic":    user.IsPublic,
		"phone":       user.Phone,
		"image":       user.Image,
	})
}
