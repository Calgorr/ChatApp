package model

import (
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
