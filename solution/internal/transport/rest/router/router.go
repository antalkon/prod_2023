package router

import (
	hmain "github.com/antalkon/prod_2023/internal/transport/rest/handler/hMain"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// NewEchoServer создает новый Echo сервер с подключенными маршрутами и middleware
func NewEchoServer() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	// Регистрация маршрутов
	RegisterRoutes(e)

	return e
}

// RegisterRoutes регистрирует маршруты для Echo сервера
func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api")
	{
		api.GET("/ping", hmain.Ping)
	}

}
