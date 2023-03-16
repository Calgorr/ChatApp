package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Calgorr/ChatApp/client/model"
)

func SendMessage(groupName string, user *model.User) error {
	var content string
	fmt.Scan(&content)
	ms := &model.Message{
		GroupName: groupName,
		Content:   content,
		Sender:    user.Username,
	}
	json, _ := json.Marshal(ms)
	req, _ := http.NewRequest("POST", "http://localhost:4545/api/messages/newmessage", bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", model.Token)
	resp, _ := http.DefaultClient.Do(req)
	if resp.StatusCode != http.StatusOK {
		return errors.New("something went wrong")
	}
	return nil
}
