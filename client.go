package main

import "github.com/gorilla/websocket"

type Client struct {
	Ws *websocket.Conn
}

// Send はクライアントにメッセージを送信する
func (client *Client) Send(msg []byte) error {
	return client.Ws.WriteMessage(websocket.TextMessage, msg)
}
