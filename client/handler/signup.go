package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Calgorr/ChatApp/client/model"
)

func SignUp(username, password string) error {
	user := &model.User{
		Username: username,
		Password: password,
	}
	credB, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "http://localhost:4545/api/users", bytes.NewBuffer(credB))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := http.DefaultClient.Do(req)
	if resp.StatusCode == http.StatusConflict {
		log.Default().Println(resp)
		return errors.New("user already exists")
	} else if resp.StatusCode != http.StatusOK {
		log.Default().Println(resp)
		return errors.New("something went wrong")
	}
	return nil
}
