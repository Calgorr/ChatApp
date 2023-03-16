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
		if err != nil {
			MainMenu()
		}
		fmt.Println(user)

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

func LoginMenu() {
	fmt.Println("1-Enter a GroupChat")
	fmt.Println("2-Create a Group")
	fmt.Println("3-Join a Group")
}
