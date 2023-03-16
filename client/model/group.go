package model

type Group struct {
	GroupName    string
	Creator      string
	CreationDate string
	Description  string
	members      []User
}
