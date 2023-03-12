package routes

import (
	handle "github.com/Calgorr/ChatApp/server/Handle"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterRoutes(g *echo.Group) {
	g.Use(middleware.AddTrailingSlash())

	user := g.Group("/users")
	user.POST("", handle.SignUp)
}
