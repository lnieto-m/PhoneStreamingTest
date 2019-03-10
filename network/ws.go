package network

import (
	"PhoneStreamingTest/adb"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  102400000000,
	WriteBufferSize: 102400000000,
}

//ServeSocket handles each websocket connexion
func ServeSocket(manager *adb.Manager, writter http.ResponseWriter, request *http.Request) {
	newConnexion, err := upgrader.Upgrade(writter, request, nil)
	if err != nil {
		log.Printf("error: %v\n", err)
		return
	}
	currentClient := &Client{
		conn: newConnexion,
		send: make(chan []byte, 256),
	}

	go currentClient.writeToSocket(manager)
	go currentClient.readSocket()
}

// ServeStatusSocket notify the IOS app when a phone status change
func ServeStatusSocket(manager *adb.Manager, writter http.ResponseWriter, request *http.Request) {
	newConnexion, err := upgrader.Upgrade(writter, request, nil)

	if err != nil {
		log.Printf("error: %v\n", err)
		return
	}
	currentClient := &Client{
		conn: newConnexion,
		send: make(chan []byte, 256),
	}
	go currentClient.notifyStatusChangeToIOS(manager)
}
