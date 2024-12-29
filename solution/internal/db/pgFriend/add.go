package pgfriend

import (
	"github.com/antalkon/prod_2023/internal/models"
	"github.com/antalkon/prod_2023/pkg/pggorm"
)

// Проверка, существует ли друг
func IsFriendAlreadyAdded(userID, friendID int) (bool, error) {
	var count int64
	err := pggorm.DB.Model(&models.Friends{}).
		Where("user_id = ? AND friend_id = ?", userID, friendID).
		Count(&count).Error
	return count > 0, err
}

// Добавление друга
func AddFriend(friend models.Friends) error {
	return pggorm.DB.Create(&friend).Error
}
