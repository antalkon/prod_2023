package hfriends

import (
	pgdata "github.com/antalkon/prod_2023/internal/db/pgData"
	pgfriend "github.com/antalkon/prod_2023/internal/db/pgFriend"
	"github.com/antalkon/prod_2023/internal/models"
	"github.com/labstack/echo/v4"
)

func Remove(c echo.Context) error {
	var fl models.AddFriend
	if err := c.Bind(&fl); err != nil {
		return err
	}

	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(401, map[string]string{
			"reason": "Переданный токен не существует либо некорректен",
		})
	}

	friend, err := pgdata.GetUserByLogin(fl.Login)
	if err != nil {
		return c.JSON(401, map[string]string{
			"reason": "Пользователь с указанным логином не найден",
		})
	}

	if err := pgfriend.RemoveFriend(userID, friend.ID); err != nil {
		return c.JSON(500, map[string]string{
			"reason": "Ошибка удаления друга",
		})
	}

	return c.JSON(200, map[string]string{
		"status": "ok",
	})

}
