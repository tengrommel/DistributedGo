package pubsub

import (
	"github.com/gorilla/websocket"
	)

type PubSub struct {
	Clients []Client
}

type Client struct {
	Id string
	Connection *websocket.Conn
}

func (ps *PubSub)AddClient(client Client)  (*PubSub) {
	ps.Clients = append(ps.Clients, client)
	//fmt.Println("adding new client to the list", client.Id, len(ps.Clients))
	payload := []byte("Hello Client ID:" +
		client.Id)
	client.Connection.WriteMessage(1, payload)
	return ps
}

func (ps *PubSub)HandleReceiveMessage(client Client, messageType int, message []byte) (*PubSub) {
	return ps
}