package network

import (
	"PhoneStreamingTest/adb"
	"fmt"
	"log"
	"strings"

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

func (client *Client) writeToSocket(manager *adb.Manager) {
	defer client.conn.Close()
	for {
		select {
		case frame := <-manager.MinicapChan:
			writter, err := client.conn.NextWriter(websocket.BinaryMessage)
			if err != nil {
				return
			}
			writter.Write(frame)
		default:
			// ????
		}
	}
	// TODO
}

func (client *Client) notifyStatusChangeToIOS(manager *adb.Manager) {
	defer client.conn.Close()
	defer func() {
		log.Print("ticker stopped")
		manager.StatusStop <- true
	}()
	for {
		select {
		case phoneID := <-manager.StatusChange:
			message := JSONMessage{
				Type:    "status",
				Message: strings.Join(phoneID, ","),
			}
			client.conn.WriteJSON(message)
			log.Println(strings.Join(phoneID, ","))
		default:
			// log.Print("No message received\n")
		}
	}
}
