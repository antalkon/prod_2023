package hauth

import (
	"net/http"

	pgdata "github.com/antalkon/prod_2023/internal/db/pgData"
	"github.com/antalkon/prod_2023/internal/models"
	"github.com/antalkon/prod_2023/pkg/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c echo.Context) error {
	var l models.Login

	// Привязка данных из тела запроса
	if err := c.Bind(&l); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"reason": "Invalid request format",
		})
	}

	// Получение пользователя из базы данных по логину
	user, err := pgdata.GetUserByLogin(l.Login)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"reason": "Invalid login or password",
		})
	}

	// Сравнение пароля
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(l.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"reason": "Invalid login or password",
		})
	}

	// Создание JWT
	token, err := jwt.CreateJWT(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"reason": "Failed to generate JWT",
		})
	}

	// Устанавливаем токен в заголовок Authorization
	c.Response().Header().Set("Authorization", "Bearer "+token)

	// Возвращаем успешный ответ
	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
