package handle

import (
	"fmt"

	db "github.com/Calgorr/ChatApp/server/database"
	"github.com/Calgorr/ChatApp/server/model"
	"github.com/labstack/echo/v4"
)

func CreateGroup(c echo.Context) error {
	//fmt.Println("skdvsdkv")
	var g *model.Group
	g, err := g.Bind(c)
	if err != nil {
		fmt.Println("skdvsd324324kv")
		return c.String(500, "internal server error")
	}
	err = db.AddGroup(g)
	if err != nil {
		fmt.Println("skdvsdkv")
		fmt.Println(err)
		return c.String(500, "internal server error")
	}
	return c.String(200, "success")
}

func AddMember(c echo.Context) error {
	var user *model.User
	user, err := user.Bind(c)
	groupName := c.QueryParam("groupname")
	if err != nil {
		return c.String(500, "internal server error")
	}
	err = db.AddMemberToGroup(groupName, user)
	if err != nil {
		return c.String(500, "internal server error")
	}
	return c.String(200, "success")
}

func GetGroups(c echo.Context) error {
	username := c.QueryParam("username")
	groups, err := db.GetGroups(username)
	if err != nil {
		return c.String(500, "internal server error")
	}
	return c.JSON(200, groups)
}

func GetMessages(c echo.Context) error {
	groupName := c.QueryParam("groupname")
	messages, err := db.GetMessages(groupName)
	if err != nil {
		return c.String(500, "internal server error")
	}
	return c.JSON(200, messages)
}
