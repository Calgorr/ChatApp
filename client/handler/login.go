package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
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
		// return fmt.Println("User does not exi")
	}
	model.Token = resp.Header.Get(echo.HeaderAuthorization)
	fmt.Println(model.Token)
	return user, nil
}

func SignUp(username, password string) error {
	user := &model.User{
		Username: username,
		Password: password,
	}
	credB, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "http://localhost:4545/api/users", bytes.NewBuffer(credB))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := http.DefaultClient.Do(req)
	if resp.StatusCode != http.StatusOK {
		log.Default().Println(resp)
		return errors.New("Internal server error")
	}
	return nil
}
