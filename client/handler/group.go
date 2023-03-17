package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/Calgorr/ChatApp/client/model"
	"github.com/labstack/echo/v4"
)

func EnterGroupChat(user *model.User, groupname string) {
	ClearConsole()
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
		fmt.Println(v.Sender, v.Send_At, v.Content)
	}

	done := make(chan bool)
	go SendMessage(groupname, user, done)
	go Receive(groupname, done)

	select {
	case <-done:
		ClearConsole()
		LoginMenu(user)
	}
}

func CreateGroup(user *model.User, groupname string) {
	var description string
	println("Enter group description: ")
	fmt.Scanln(&description)

	group := &model.Group{
		GroupName:    groupname,
		Description:  description,
		Creator:      user.Username,
		CreationDate: time.Now(),
	}
	json, _ := json.Marshal(group)
	req, _ := http.NewRequest(http.MethodPost, "localhost:4545/api/groups/newgroup", bytes.NewBuffer(json))
	req.Header.Set(echo.HeaderAuthorization, model.Token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		println("Group created successfully")
	} else {
		println("Group creation failed")
	}
	LoginMenu(user)

}

func JoinGroup(user *model.User) {
	var groupname string
	println("Enter group name: ")
	fmt.Scanln(&groupname)

	req, _ := http.NewRequest(http.MethodGet, "localhost:4545/api/groups/joingroup?groupname="+groupname, nil)
	req.Header.Set(echo.HeaderAuthorization, model.Token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		println("Joined group successfully")
	} else {
		println("Group join failed")
	}
	LoginMenu(user)
}

func LoginMenu(user *model.User) {
	fmt.Println("1-Enter a GroupChat")
	fmt.Println("2-Create a Group")
	fmt.Println("3-Join a Group")
	var order int
	fmt.Scanln(&order)
	switch order {
	case 1:
		var groupname string
		fmt.Println("Group name :")
		fmt.Scan(&groupname)
		EnterGroupChat(user, groupname)
	case 2:
		var groupName string
		fmt.Println("Group name :")
		fmt.Scan(&groupName)
		CreateGroup(user, groupName)
	case 3:
		JoinGroup(user)

	}

}

func ClearConsole() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
