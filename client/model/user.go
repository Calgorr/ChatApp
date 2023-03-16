package model

type User struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

var Token string
