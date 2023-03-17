package model

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Group struct {
	GroupName    string `json:"groupname" form:"groupname"`
	Creator      string `json:"creator" form:"creator"`
	CreationDate string `json:"creationdate" form:"creationdate"`
	Description  string `json:"description" form:"description"`
	Members      []User `json:"members" form:"members"`
}

func (g *Group) AddMember(u User) {
	g.Members = append(g.Members, u)
}

func (g *Group) CheckMember(u User) bool {
	for _, member := range g.Members {
		if member.Username == u.Username {
			return true
		}
	}
	return false
}

func (g *Group) RemoveMember(username string) {
	for i, member := range g.Members {
		if member.Username == username {
			g.Members = append(g.Members[:i], g.Members[i+1:]...)
		}
	}
}

func (g *Group) Bind(c echo.Context) (*Group, error) {
	err := c.Bind(&g)
	if err != nil {
		return nil, c.String(http.DefaultMaxHeaderBytes, "bad request")
	}
	return g, nil
}

func (g *Group) MarshalBinary() ([]byte, error) {
	return json.Marshal(g)
}

func (g *Group) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, g)
}
