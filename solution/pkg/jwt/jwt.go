package jwt

import (
	"errors"
	"time"

	"github.com/antalkon/prod_2023/internal/config"
	"github.com/golang-jwt/jwt/v4"
)

// Claims структура, содержащая данные JWT
type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

// CreateJWT создает JWT с `user_id` и возвращает токен и ошибку
func CreateJWT(userID int) (string, error) {
	// Получаем конфигурацию
	cfg := config.GetConfig()

	// Создаем claims
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // JWT будет действителен 24 часа
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Создаем токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен с использованием секретного ключа
	tokenString, err := token.SignedString([]byte(cfg.RandomSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// DecodeJWT принимает JWT, декодирует его и возвращает user_id и ошибку
func DecodeJWT(tokenString string) (int, error) {
	// Получаем конфигурацию
	cfg := config.GetConfig()

	// Парсинг токена
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Проверяем метод подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(cfg.RandomSecret), nil
	})
	if err != nil {
		return 0, err
	}

	// Извлечение данных из claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.UserID, nil
	}

	return 0, errors.New("invalid token")
}
