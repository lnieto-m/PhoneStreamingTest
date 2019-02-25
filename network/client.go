package network

import (
	"PhoneStreamingTest/adb"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

// Client : object holding current client connexion
type Client struct {
	conn *websocket.Conn

	send chan []byte
}

// JSONMessage : base struct for json message for this server
type JSONMessage struct {
	Type    string `json:"type"`
	Message string `json:"message"`
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

func (client *Client) notifyStatusChangeToIOS(manager *adb.Manager) {
	defer client.conn.Close()
	for {
		select {
		case phoneID := <-manager.StatusChange:
			message := JSONMessage{
				Type:    "status",
				Message: phoneID,
			}
			client.conn.WriteJSON(message)
		default:
			// log.Print("No message received\n")
		}
	}
}
