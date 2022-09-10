package main

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"net/http"
)

var ws *websocket.Conn

func Hello(c echo.Context) error {
	// Write
	err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
	if err != nil {
		c.Logger().Error(err)
	}
	return c.String(http.StatusOK, "Hello, World!")
}
