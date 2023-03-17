package handle

import (
	"fmt"

	db "github.com/Calgorr/ChatApp/server/database"
	"github.com/Calgorr/ChatApp/server/model"
	"github.com/labstack/echo/v4"
)

func CreateGroup(c echo.Context) error {
	var g *model.Group
	g, err := g.Bind(c)
	if err != nil {
		return c.String(500, "internal server error")
	}
	err = db.AddGroup(g)
	if err != nil {
		return c.String(500, "internal server error")
	}
	return c.String(200, "success")
}

func AddMember(c echo.Context) error {
	fmt.Println("skfnvjfvnj")
	var user *model.User
	user, err := user.Bind(c)
	groupName := c.QueryParam("groupname")
	fmt.Println(groupName)
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
	if db.CheckGroup(c.QueryParam("groupname")) == false {
		return c.String(404, "group not found")
	}
	groupName := c.QueryParam("groupname")
	messages, err := db.GetMessages(groupName)
	if err != nil {
		return c.String(500, "internal server error")
	}
	fmt.Println(messages)
	return c.JSON(200, messages)
}

func GetGroup(c echo.Context) error {
	groupName := c.QueryParam("groupname")
	group, err := db.GetGroup(groupName)
	for _, v := range group.Members {
		v.Password = ""
	}
	if err != nil {
		return c.String(500, "internal server error")
	}
	return c.JSON(200, group)
}

func RemoveUser(c echo.Context) error {
	groupName := c.QueryParam("groupname")
	username := c.QueryParam("username")
	err := db.RemoveUserFromGroup(groupName, username)
	if err != nil {
		fmt.Println(err)
		return c.String(500, "internal server error")
	}
	return c.String(200, "success")
}
