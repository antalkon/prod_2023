package pgauth

import (
	"fmt"

	"github.com/antalkon/prod_2023/internal/models"
	"github.com/antalkon/prod_2023/pkg/pggorm"
)

func CreateUser(u models.User) (models.User, error) {
	db := pggorm.DB
	if db == nil {
		return models.User{}, fmt.Errorf("database connection is not initialized")
	}

	// Проверка на уникальность email, логина или телефона
	var existingUser models.User
	if err := db.Where("email = ? OR phone = ? OR login = ?", u.Email, u.Phone, u.Login).First(&existingUser).Error; err == nil {
		return models.User{}, fmt.Errorf("user with provided email, phone, or login already exists")
	}

	// Сохранение нового пользователя
	if err := db.Create(&u).Error; err != nil {
		return models.User{}, fmt.Errorf("failed to create user, error: %w", err)
	}

	return u, nil
}
