package main

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"net/http"
)

var upgrader = websocket.Upgrader{}
var rooms = Room{}

func ServeWs(c echo.Context) error {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		c.Logger().Error(err)
	}
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			c.Logger().Error(err)
		}
	}(ws)

	client := &Client{Ws: ws}

	rooms.AddClient(client)
	for {
		//
		// _, msg, err := ws.ReadMessage()
		// if err != nil {
		// 	c.Logger().Error(err)
		// 	break
		// }
		// ReadMessage() によって受け取ったメッセージを Publish() で配信する
		msg := <-C1
		rooms.Publish([]byte(msg))
	}

	// return nil
}
