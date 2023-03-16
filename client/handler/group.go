package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Calgorr/ChatApp/client/cli"
	"github.com/Calgorr/ChatApp/client/model"
	"github.com/labstack/echo/v4"
)

func EnterGroupChat(user *model.User, groupname string) {
	cli.ClearConsole()
	var messages []model.Message

	req, _ := http.NewRequest(http.MethodGet, "localhost:4545/api/groups/getmessages?groupname="+groupname, nil)
	req.Header.Set(echo.HeaderAuthorization, model.Token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&messages)
	if err != nil {
		panic(err)
	}
	for _, v := range messages {
		println(v.Sender, v.Send_At, v.Content)
	}

	go SendMessage(groupname, user)
	go Receive(groupname)

}

func CreateGroup(user *model.User, groupname string) {
	// TODO
}

func JoinGroup(user *model.User) {
	// TODO
}
