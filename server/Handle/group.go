package handle

import (
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
	var user *model.User
	user, err := user.Bind(c)
	groupName:=c.QueryParam("groupname")
	if err != nil {
		return c.String(500, "internal server error")
	}
	err = db.AddMemberToGroup(groupName,user.Username)
	if err != nil {

}
