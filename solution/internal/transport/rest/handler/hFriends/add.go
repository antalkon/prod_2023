package hfriends

import (
	"net/http"

	pgdata "github.com/antalkon/prod_2023/internal/db/pgData"
	pgfriend "github.com/antalkon/prod_2023/internal/db/pgFriend"
	"github.com/antalkon/prod_2023/internal/models"
	"github.com/labstack/echo/v4"
)

func Add(c echo.Context) error {
	// Получаем идентификатор пользователя из контекста
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"reason": "Переданный токен не существует либо некорректен",
		})
	}

	// Парсим тело запроса в структуру
	var addLogin struct {
		Login string `json:"login"`
	}
	if err := c.Bind(&addLogin); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"reason": "Некорректный запрос, проверьте переданные данные",
		})
	}

	// Ищем пользователя по логину
	frID, err := pgdata.GetUserByLogin(addLogin.Login)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"reason": "Пользователь с указанным логином не найден",
		})
	}

	// Проверяем, не пытается ли пользователь добавить сам себя
	if frID.ID == userID {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	}

	// Проверяем, не добавлен ли уже друг
	exists, err := pgfriend.IsFriendAlreadyAdded(userID, frID.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"reason": "Ошибка проверки существующих друзей",
		})
	}
	if exists {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	}

	// Добавляем друга
	friend := models.Friends{
		UserID:   userID,
		FriendID: frID.ID,
	}
	if err := pgfriend.AddFriend(friend); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"reason": "Ошибка добавления друга в базу данных",
		})
	}

	// Успешный ответ
	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}
