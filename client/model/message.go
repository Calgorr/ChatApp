package model

import (
	"encoding/json"
	"time"
)

type Message struct {
	Send_At   time.Time `json:"date" form:"date"`
	Sender    string    `json:"sender" form:"sender"`
	GroupName string    `json:"groupname" form:"groupname"`
	Content   string    `json:"content" form:"content"`
}

func (ms *Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(ms)
}

func (ms *Message) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, ms)
}
