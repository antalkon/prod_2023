package pgfriend

import (
	"fmt"

	"github.com/antalkon/prod_2023/internal/models"
	"github.com/antalkon/prod_2023/pkg/pggorm"
)

func GetAllUserFriends(userID, limit, offset int) ([]models.FriendResponse, error) {
	db := pggorm.DB
	if db == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	var friends []models.FriendResponse
	query := `
		SELECT f.friend_id, f.added_at
		FROM friends f
		WHERE f.user_id = ?
		ORDER BY f.added_at DESC
		LIMIT ? OFFSET ?`

	if err := db.Raw(query, userID, limit, offset).Scan(&friends).Error; err != nil {
		return nil, fmt.Errorf("failed to get friends for user %d: %w", userID, err)
	}

	return friends, nil
}
