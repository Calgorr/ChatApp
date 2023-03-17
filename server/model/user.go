package model

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (user *User) Bind(c echo.Context) (*User, error) {
	err := c.Bind(&user)
	if err != nil {
		fmt.Println(err)
		return nil, c.String(http.DefaultMaxHeaderBytes, "bad request")
	}
	return user, nil
}
