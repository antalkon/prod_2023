package pgfriend

import (
	"fmt"

	"github.com/antalkon/prod_2023/internal/models"
	"github.com/antalkon/prod_2023/pkg/pggorm"
)

func RemoveFriend(userID, friendID int) error {
	db := pggorm.DB
	if db == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	return db.Where("user_id = ? AND friend_id =?", userID, friendID).Delete(&models.Friends{}).Error
}
