package model

import "time"

type Message struct {
	Send_At   time.Time `json:"date" form:"date"`
	Sender    string    `json:"sender" form:"sender"`
	GroupName string    `json:"groupname" form:"groupname"`
	Content   string    `json:"content" form:"content"`
}
