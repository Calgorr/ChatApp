package cli

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Calgorr/ChatApp/client/handler"
	"github.com/Calgorr/ChatApp/client/model"
)

func MainMenu() {
	fmt.Println("Welcome to the chat app!")
	fmt.Println("1. Login")
	fmt.Println("2. Sign up")
	var order int
	fmt.Scanln(&order)
	switch order {
	case 1:
		var username, password string
		fmt.Println("Username :")
		fmt.Scan(&username)
		fmt.Println("Password :")
		fmt.Scan(&password)
		user, err := handler.Login(username, password)
		if err != nil {
			MainMenu()
		}
		LoginMenu(user)

	case 2:
		var username, password string
		fmt.Println("Username :")
		fmt.Scan(&username)
		fmt.Println("Password :")
		fmt.Scan(&password)
		handler.SignUp(username, password)
		MainMenu()
	}
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
		handler.EnterGroupChat(user, groupname)
	case 2:
		var groupName string
		fmt.Println("Group name :")
		fmt.Scan(&groupName)
		handler.CreateGroup(user, groupName)
	case 3:
		handler.JoinGroup(user)

	}

}

func ClearConsole() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
