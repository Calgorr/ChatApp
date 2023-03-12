package handle

import (
	"encoding/json"
	"net/http"

	authentication "github.com/Calgorr/ChatApp/server/Authentication"
	db "github.com/Calgorr/ChatApp/server/database"
	"github.com/Calgorr/ChatApp/server/model"
	"github.com/labstack/echo/v4"
)

func SignUp(c echo.Context) error {
	var user *model.User
	user, err := bind(c, user)
	//add to database
	c.String(http.StatusOK, "success")
	return err
}

func Login(c echo.Context) error {
	var user *model.User
	user, err := bind(c, user)
	//check database
	token, err := authentication.GenerateJWT()
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	c.Response().Header().Set(echo.HeaderAuthorization, token)

	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(user)
}

func userValidation(user *model.User) bool {
	user, err := db.GetUser(user.Username, user.Password)
	if err != nil {
		return false
	}
	return true
}

func bind(c echo.Context, user *model.User) (*model.User, error) {
	err := c.Bind(&user)
	if err != nil {
		return nil, c.String(http.DefaultMaxHeaderBytes, "bad request")
	}
	return user, nil
}
