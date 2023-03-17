package model

import "time"

type Group struct {
	GroupName    string
	Creator      string
	CreationDate time.Time
	Description  string
	Members      []User
}
