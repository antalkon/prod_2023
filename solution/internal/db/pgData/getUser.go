package pgdata

import (
	"errors"
	"fmt"

	"github.com/antalkon/prod_2023/internal/models"
	"github.com/antalkon/prod_2023/pkg/pggorm"
	"gorm.io/gorm"
)

// GetUserByLogin ищет пользователя по логину
func GetUserByLogin(login string) (models.User, error) {
	db := pggorm.DB
	if db == nil {
		return models.User{}, fmt.Errorf("database connection is not initialized")
	}

	var user models.User
	err := db.Where("login = ?", login).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.User{}, fmt.Errorf("user not found")
	}
	if err != nil {
		return models.User{}, fmt.Errorf("error querying database: %w", err)
	}

	return user, nil
}

func GetUserById(id string) (models.User, error) {
	db := pggorm.DB
	if db == nil {
		return models.User{}, fmt.Errorf("database connection is not initialized")
	}

	var user models.User
	err := db.Where("id =?", id).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.User{}, fmt.Errorf("user not found")
	}
	if err != nil {
		return models.User{}, fmt.Errorf("error querying database: %w", err)
	}

	return user, nil
}
