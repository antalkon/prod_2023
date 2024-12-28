package hauth

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	pgauth "github.com/antalkon/prod_2023/internal/db/pgAuth"
	pgmain "github.com/antalkon/prod_2023/internal/db/pgMain"
	"github.com/antalkon/prod_2023/internal/models"
	"github.com/antalkon/prod_2023/pkg/hash"
	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	var u models.User

	// Привязка данных из тела запроса
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"reason": "Invalid request format",
		})
	}

	// Валидация входных данных
	if err := validateUserInput(u); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"reason": err.Error(),
		})
	}

	// Проверка существования страны по коду
	country, err := pgmain.GetCountryByAlpha2(u.CountryCode)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"reason": "Country with the provided code does not exist",
		})
	}
	fmt.Printf("Country validated: %s\n", country.Name)

	// Генерация хэша пароля
	hash, err := hash.HashPassword(u.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"reason": "Failed to hash password",
		})
	}
	u.PasswordHash = hash
	u.Password = "" // Убираем исходный пароль

	// Создание пользователя в базе данных
	createdUser, err := pgauth.CreateUser(u)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			return c.JSON(http.StatusConflict, map[string]string{
				"reason": "User with provided email, phone, or login already exists",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"reason": "Failed to create user",
		})
	}

	// Формирование успешного ответа
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"profile": map[string]interface{}{
			"login":       createdUser.Login,
			"email":       createdUser.Email,
			"countryCode": createdUser.CountryCode,
			"isPublic":    createdUser.IsPublic,
			"phone":       createdUser.Phone,
			"image":       createdUser.Image,
		},
	})
}

// validateUserInput проверяет корректность входных данных
func validateUserInput(u models.User) error {
	if len(u.Login) > 30 || !regexp.MustCompile(`^[a-zA-Z0-9-]+$`).MatchString(u.Login) {
		return fmt.Errorf("Invalid login: must be alphanumeric, may contain dashes, and not exceed 30 characters")
	}

	if len(u.Email) > 50 || !regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`).MatchString(u.Email) {
		return fmt.Errorf("Invalid email format")
	}

	if len(u.Password) < 8 {
		return fmt.Errorf("Password is too weak: must be at least 8 characters long")
	}

	if len(u.CountryCode) != 2 || !regexp.MustCompile(`^[a-zA-Z]{2}$`).MatchString(u.CountryCode) {
		return fmt.Errorf("Invalid country code: must be exactly 2 alphabetic characters")
	}

	if u.Phone != "" && !regexp.MustCompile(`^\+\d+$`).MatchString(u.Phone) {
		return fmt.Errorf("Invalid phone: must start with + and contain only digits")
	}

	if len(u.Image) > 200 {
		return fmt.Errorf("Image URL exceeds maximum length of 200 characters")
	}

	return nil
}
