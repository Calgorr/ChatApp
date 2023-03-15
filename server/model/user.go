package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (user *User) ValidatePassword(pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)) == nil
}
