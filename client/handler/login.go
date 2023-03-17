package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Calgorr/ChatApp/client/model"
	"github.com/labstack/echo/v4"
)

func Login(username, password string) (*model.User, error) {
	user := &model.User{
		Username: username,
		Password: password,
	}
	credB, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "http://localhost:4545/api/users/login", bytes.NewBuffer(credB))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := http.DefaultClient.Do(req)
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Login failed")
		return nil, fmt.Errorf("login failed")
	}
	model.Token = resp.Header.Get(echo.HeaderAuthorization)
	return user, nil
}
