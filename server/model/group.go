package model

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Group struct {
	GroupName    string
	Creator      string
	CreationDate string
	Description  string
	members      []User
}

func (g *Group) AddMember(u User) {
	g.members = append(g.members, u)
}

func (g *Group) CheckMember(u User) bool {
	for _, member := range g.members {
		if member.Username == u.Username {
			return true
		}
	}
	return false
}

func (g *Group) Bind(c echo.Context) (*Group, error) {
	err := c.Bind(&g)
	if err != nil {
		return nil, c.String(http.DefaultMaxHeaderBytes, "bad request")
	}
	return g, nil
}
