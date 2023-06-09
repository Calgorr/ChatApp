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

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:4545/api/groups/getmessages?groupname="+groupname, nil)
	req.Header.Set(echo.HeaderAuthorization, model.Token)
	res, err := http.DefaultClient.Do(req)
	if res.StatusCode == http.StatusNotFound {
		fmt.Println("Group not found")
		time.Sleep(2 * time.Second)
		ClearConsole()
		LoginMenu(user)
	}

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&messages)
	if err != nil {
		panic(err)
	}
	for _, v := range messages {
		fmt.Println(v.Sender, ": ", v.Content)
	}

	done := make(chan bool)
	go SendMessage(groupname, user, done)
	go Receive(groupname, user.Username, done)

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
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:4545/api/groups/newgroup", bytes.NewBuffer(json))
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
	ClearConsole()
	LoginMenu(user)

}

func JoinGroup(user *model.User) {
	var groupname string
	println("Enter group name: ")
	fmt.Scanln(&groupname)
	json, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:4545/api/groups/addmember?groupname="+groupname, bytes.NewBuffer(json))
	req.Header.Set(echo.HeaderAuthorization, model.Token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		JoinGroup(user)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		println("Joined group successfully")
	} else {
		println("Group join failed")
	}
	ClearConsole()
	LoginMenu(user)
}

func GetGroupInfo(groupname string) {
	var group model.Group
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:4545/api/groups/getgroup?groupname="+groupname, nil)
	req.Header.Set(echo.HeaderAuthorization, model.Token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&group)
	if err != nil {
		panic(err)
	}
	fmt.Println("Group name: ", group.GroupName)
	fmt.Println("Group description: ", group.Description)
	fmt.Println("Group creator: ", group.Creator)
	fmt.Println("Group creation date: ", group.CreationDate)
	fmt.Println("Group members: ")
	for _, v := range group.Members {
		fmt.Println(v.Username)
	}
}

func RemoveUserFromGroup(groupname string, user *model.User) {
	req, _ := http.NewRequest(http.MethodDelete, "http://localhost:4545/api/groups/removeuser?groupname="+groupname+"&username="+user.Username, nil)
	req.Header.Set(echo.HeaderAuthorization, model.Token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	ClearConsole()
	if res.StatusCode == 200 {
		println("User removed successfully")
	} else {
		println("User remove failed")
	}
	LoginMenu(user)
}

func LoginMenu(user *model.User) {
	fmt.Println("1-Enter a GroupChat")
	fmt.Println("2-Create a Group")
	fmt.Println("3-Join a Group")
	fmt.Println("4-Remove from a Group")
	var order int
	fmt.Scanln(&order)
	switch order {
	case 1:
		var groupname string
		fmt.Println("Group name :")
		fmt.Scan(&groupname)
		EnterGroupChat(user, groupname)
		ClearConsole()
	case 2:
		var groupName string
		fmt.Println("Group name :")
		fmt.Scan(&groupName)
		CreateGroup(user, groupName)
		ClearConsole()
	case 3:
		JoinGroup(user)
		ClearConsole()
	case 4:
		var groupname string
		fmt.Println("Group name :")
		fmt.Scan(&groupname)
		RemoveUserFromGroup(groupname, user)
	}

}

func ClearConsole() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
