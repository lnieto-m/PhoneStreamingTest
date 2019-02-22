package network

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

// Client : object holding current client connexion
type Client struct {
	conn *websocket.Conn

	send chan []byte
}

func (client *Client) readSocket() {
	defer client.conn.Close()
	for {
		_, message, err := client.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		fmt.Println(string(message))
		// TODO: Command treatment when VM is setup
	}
}

func (client *Client) writeToSocket() {
	defer client.conn.Close()
	// TODO
}
