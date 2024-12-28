package pgdata

import (
	"errors"
	"fmt"
	"strings"

	"github.com/antalkon/prod_2023/internal/models"
	"github.com/antalkon/prod_2023/pkg/pggorm"
)

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrDuplicateEntry = errors.New("duplicate entry")
)

func UpdateUserProfile(user models.User) error {
	db := pggorm.DB
	if db == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	// Сохранение данных пользователя
	err := db.Save(&user).Error
	if err != nil {
		if isDuplicateError(err) {
			return ErrDuplicateEntry
		}
		return fmt.Errorf("failed to update user profile: %w", err)
	}

	return nil
}

// isDuplicateError проверяет ошибку на уникальность
func isDuplicateError(err error) bool {
	return strings.Contains(err.Error(), "duplicate key value")
}
