package router

import (
	middleware1 "github.com/antalkon/prod_2023/internal/middleware"
	hauth "github.com/antalkon/prod_2023/internal/transport/rest/handler/hAuth"
	hfriends "github.com/antalkon/prod_2023/internal/transport/rest/handler/hFriends"
	hmain "github.com/antalkon/prod_2023/internal/transport/rest/handler/hMain"
	hme "github.com/antalkon/prod_2023/internal/transport/rest/handler/hMe"
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
		api.GET("/countries", hmain.Countries)
		api.GET("/countries/:alpha2", hmain.GetCountry)
	}

	auth := api.Group("/auth")
	{
		auth.POST("/register", hauth.Register)
		auth.POST("/sign-in", hauth.Login)

	}

	me := api.Group("/me", middleware1.AuthMiddleware)
	{
		me.GET("/profile", hme.MyProfile)
		me.PATCH("/profile", hme.EditMyProfile)
		me.POST("/updatePassword", hme.UpdPsw)

	}
	friends := api.Group("/friends", middleware1.AuthMiddleware)

	{
		friends.GET("", hfriends.Friends)

		friends.POST("/add", hfriends.Add)
		friends.POST("/remove", hfriends.Remove)

	}

}
