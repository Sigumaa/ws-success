package main

import "github.com/gorilla/websocket"

type Client struct {
	Ws *websocket.Conn
}

// Send はクライアントにメッセージを送信する
func (client *Client) Send(msg []byte) error {

	// メッセージを送信するにはWriteMessageを使う
	return client.Ws.WriteMessage(websocket.TextMessage, msg)
}

// SendJSON はクライアントにJSONを送信する
func (client *Client) SendJSON(msg any) error {

	// JSONを送信するにはWriteJSONを使う
	return client.Ws.WriteJSON(msg)
}
