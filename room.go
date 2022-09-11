package main

type Room struct {
	Clients []*Client
}

//	type Client struct {
//	   Ws *websocket.Conn
//	}
//
// add new client to room
func (room *Room) AddClient(client *Client) {
	// type Room struct {
	//	Clients []*Client
	// }
	// append client to room
	room.Clients = append(room.Clients, client)
}

// PublishMSG メッセージを接続しているクライアントに送信する
func (room *Room) PublishMSG(msg []byte) {
	for _, client := range room.Clients {
		err := client.Send(msg)
		if err != nil {
			return
		}
	}
}

// PublishJSON JSONを接続しているクライアントに送信する
func (room *Room) PublishJSON(json any) {
	for _, client := range room.Clients {
		err := client.SendJSON(json)
		if err != nil {
			return
		}
	}
}
