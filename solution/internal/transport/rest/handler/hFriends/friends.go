package hfriends

import (
	"net/http"
	"strconv"

	pgdata "github.com/antalkon/prod_2023/internal/db/pgData"
	pgfriend "github.com/antalkon/prod_2023/internal/db/pgFriend"
	"github.com/labstack/echo/v4"
)

func Friends(c echo.Context) error {
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"reason": "Переданный токен не существует либо некорректен",
		})
	}

	// Чтение параметров пагинации
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")

	// Установка значений по умолчанию
	limitVal := 5
	offsetVal := 0

	if limit != "" {
		if l, err := strconv.Atoi(limit); err == nil {
			limitVal = l
		}
	}

	if offset != "" {
		if o, err := strconv.Atoi(offset); err == nil {
			offsetVal = o
		}
	}

	// Получение списка друзей
	friends, err := pgfriend.GetAllUserFriends(userID, limitVal, offsetVal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"reason": "Ошибка получения списка друзей",
		})
	}

	// Форматирование результата
	response := []map[string]interface{}{}
	for _, friend := range friends {
		// Получение логина друга по его ID
		loginData, err := pgdata.GetUserById(strconv.Itoa(friend.FriendID))
		if err != nil {
			continue // Пропускаем ошибку, если не удается получить логин
		}
		response = append(response, map[string]interface{}{
			"login":   loginData.Login,
			"addedAt": friend.AddedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return c.JSON(http.StatusOK, response)
}
