package main

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	upgrader = websocket.Upgrader{}
)

func hello(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			c.Logger().Error(err)
		}
	}(ws)

	// Write
	err = ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
	if err != nil {
		c.Logger().Error(err)
	}
	return nil
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "./")
	e.GET("/ws", hello)
	e.Logger.Fatal(e.Start(":1323"))
}
