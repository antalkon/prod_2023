package hme

import (
	"net/http"
	"regexp"
	"strconv"

	pgdata "github.com/antalkon/prod_2023/internal/db/pgData"
	"github.com/antalkon/prod_2023/pkg/hash"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// validatePassword проверяет соответствие пароля требованиям безопасности
func validatePassword(password string) error {
	if len(password) < 6 || len(password) > 100 {
		return echo.NewHTTPError(http.StatusBadRequest, "Password must be between 6 and 100 characters")
	}
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return echo.NewHTTPError(http.StatusBadRequest, "Password must contain at least one uppercase letter")
	}
	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return echo.NewHTTPError(http.StatusBadRequest, "Password must contain at least one lowercase letter")
	}
	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return echo.NewHTTPError(http.StatusBadRequest, "Password must contain at least one digit")
	}
	return nil
}

func UpdPsw(c echo.Context) error {
	var psw struct {
		OldPassword string `json:"oldPassword"`
		NewPassword string `json:"newPassword"`
	}
	if err := c.Bind(&psw); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"reason": "Invalid request format",
		})
	}

	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"reason": "Unauthorized: missing or invalid user ID",
		})
	}

	user, err := pgdata.GetUserById(strconv.Itoa(userID))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"reason": "Invalid token or user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(psw.OldPassword)); err != nil {
		return c.JSON(http.StatusForbidden, map[string]string{
			"reason": "Provided old password does not match the current password",
		})
	}

	if err := validatePassword(psw.NewPassword); err != nil {
		return err
	}

	newPasswordHash, err := hash.HashPassword(psw.NewPassword)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"reason": "Failed to hash new password",
		})
	}

	if err := pgdata.UpdateUserPassword(userID, newPasswordHash); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"reason": "Failed to update password",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}
