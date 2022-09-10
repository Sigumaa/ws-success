package main

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{}
)

func Connect(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	// Write
	err = ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
	if err != nil {
		c.Logger().Error(err)
	}
	return nil
}

//func Send(ws *websocket.Conn) error {
//}
