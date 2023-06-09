package cli

import (
	"fmt"

	"github.com/Calgorr/ChatApp/client/handler"
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
		handler.ClearConsole()
		if err != nil {
			fmt.Println(err)
			MainMenu()
		}
		handler.LoginMenu(user)

	case 2:
		var username, password string
		fmt.Println("Username :")
		fmt.Scan(&username)
		fmt.Println("Password :")
		fmt.Scan(&password)
		err := handler.SignUp(username, password)
		handler.ClearConsole()
		if err != nil {
			fmt.Println(err)
		}
		MainMenu()
	}
}
