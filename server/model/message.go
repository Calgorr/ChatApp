package model

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Message struct {
	Send_At   time.Time `json:"date" form:"date"`
	Sender    string    `json:"sender" form:"sender"`
	GroupName string    `json:"groupname" form:"groupname"`
	Content   string    `json:"content" form:"content"`
}

func (ms *Message) Bind(c echo.Context) (*Message, error) {
	err := c.Bind(&ms)
	if err != nil {
		return nil, c.String(http.DefaultMaxHeaderBytes, "bad request")
	}
	return ms, nil
}

func (ms *Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(ms)
}

func (ms *Message) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, ms)
}
