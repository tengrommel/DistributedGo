package main

import (
	"net/http"
	"github.com/gorilla/websocket"
	"log"
	"fmt"
	"github.com/satori/go.uuid"
	"DistributedGo/ch14/pubsub"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func autoId() (string) {
	return uuid.Must(uuid.NewV4(), nil).String()
}

var ps = &pubsub.PubSub{}

func websocketHandler(w http.ResponseWriter, r *http.Request)  {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil{
		log.Println(err)
		return
	}

	fmt.Println("New Client is connected")

	client := pubsub.Client{
		Id:autoId(),
		Connection: conn,
	}

	ps.AddClient(client)

	for{
		messageType, p, err := conn.ReadMessage()
		if err != nil{
			log.Println(err)
			return
		}
		aMessage := []byte("Hi Client i'm server")
		if err := conn.WriteMessage(messageType, aMessage); err != nil{
			log.Println(err)
			return
		}
		fmt.Printf("New message from Client:%s", p)
	}
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "static")
	})
	http.HandleFunc("/ws",  websocketHandler)
	http.ListenAndServe(":3000", nil)
	fmt.Println("Server is running: http://localhost:3000")
}
