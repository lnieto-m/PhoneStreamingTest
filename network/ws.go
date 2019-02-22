package network

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//ServeSocket handles each websocket connexion
func ServeSocket(writter http.ResponseWriter, request *http.Request) {
	newConnexion, err := upgrader.Upgrade(writter, request, nil)
	if err != nil {
		log.Printf("error: %v\n", err)
		return
	}
	currentClient := &Client{
		conn: newConnexion,
		send: make(chan []byte, 256),
	}

	go currentClient.readSocket()
}
