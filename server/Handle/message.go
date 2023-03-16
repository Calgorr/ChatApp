package handle

import (
	"time"

	db "github.com/Calgorr/ChatApp/server/database"
	"github.com/Calgorr/ChatApp/server/model"
	"github.com/labstack/echo/v4"
)

func SendMessage(c echo.Context) error {
	var ms *model.Message
	ms, err := ms.Bind(c)
	ms.Send_At = time.Now()
	err = db.AddMessage(ms)
	err = db.Publish(ms)
	if err != nil {
		return c.String(500, "internal server error")
	}
	return c.String(200, "success")
}
