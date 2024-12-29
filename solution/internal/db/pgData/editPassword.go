package pgdata

import (
	"fmt"

	"github.com/antalkon/prod_2023/internal/models"
	"github.com/antalkon/prod_2023/pkg/pggorm"
)

func UpdateUserPassword(userID int, newPasswordHash string) error {
	db := pggorm.DB
	if db == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	// Обновление пароля пользователя
	result := db.Model(&models.User{}).Where("id = ?", userID).Update("password_hash", newPasswordHash)
	if result.Error != nil {
		return fmt.Errorf("failed to update user password: %w", result.Error)
	}

	// Проверка, обновлена ли запись
	if result.RowsAffected == 0 {
		return ErrUserNotFound
	}

	return nil
}
