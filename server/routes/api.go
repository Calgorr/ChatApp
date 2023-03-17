package routes

import (
	authentication "github.com/Calgorr/ChatApp/server/Authentication"
	handle "github.com/Calgorr/ChatApp/server/Handle"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterRoutes(g *echo.Group) {
	g.Use(middleware.Logger())
	user := g.Group("/users")
	user.POST("", handle.SignUp)
	user.POST("/login", handle.Login)

	message := g.Group("/messages")
	message.Use(authentication.ValidateJWT)
	message.POST("/newmessage", handle.SendMessage)

	groups := g.Group("/groups")
	groups.Use(authentication.ValidateJWT)
	groups.POST("/newgroup", handle.CreateGroup)
	groups.POST("/addmember", handle.AddMember)
	groups.POST("/getgroups", handle.GetGroups)
	groups.GET("/getmessages", handle.GetMessages)
	groups.GET("/getgroup", handle.GetGroup)

}
