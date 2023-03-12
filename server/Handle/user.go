package handle

import (
	"net/http"

	"github.com/Calgorr/ChatApp/server/model"
	"github.com/labstack/echo/v4"
)

func SignUp(c echo.Context) error {
	var user *model.User
	user, err := bind(c, user)
	//add to database
	return err
}

func bind(c echo.Context, user *model.User) (*model.User, error) {
	err := c.Bind(&user)
	if err != nil {
		return nil, c.String(http.DefaultMaxHeaderBytes, "bad request")
	}
	return user, nil
}
