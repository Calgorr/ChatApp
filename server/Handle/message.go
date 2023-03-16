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
	if err != nil {
		return err
	}
	ms.Send_At = time.Now()
	db.AddMessage(ms)
	db.Publish(ms)
	return nil
}
