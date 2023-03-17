package handler

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/Calgorr/ChatApp/client/model"
)

func SendMessage(groupName string, user *model.User, signal chan<- bool) error {
	for {
		var content string
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		content = scanner.Text()
		if content == "--quit" {
			signal <- true
			break
		} else if content == "--information" {
			GetGroupInfo(groupName)
			continue
		}
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
	}
	return nil
}
