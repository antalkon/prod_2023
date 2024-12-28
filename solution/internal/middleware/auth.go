package middleware1

import (
	"net/http"
	"strings"

	"github.com/antalkon/prod_2023/pkg/jwt"
	"github.com/labstack/echo/v4"
)

// AuthMiddleware проверяет наличие и валидность JWT в заголовке Authorization
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Получаем заголовок Authorization
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Authorization header is missing",
			})
		}

		// Проверяем формат заголовка
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Invalid Authorization header format",
			})
		}

		// Декодируем токен
		token := parts[1]
		userID, err := jwt.DecodeJWT(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Invalid or expired token",
			})
		}

		// Устанавливаем user_id в контекст, чтобы использовать в хендлерах
		c.Set("user_id", userID)

		// Передаем управление следующему хендлеру
		return next(c)
	}
}
