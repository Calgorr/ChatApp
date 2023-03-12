package main

import (
	"github.com/Calgorr/ChatApp/server/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	routes.RegisterRoutes(e.Group("/api"))
	e.Logger.Fatal(e.Start(":8080"))
}
