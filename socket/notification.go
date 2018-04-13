package socket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type notification struct {
	*websocket.Conn
	Hub chan string
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  0,
	WriteBufferSize: 1024,
}

// Notices globally holds user msg that we want to send, the Push func is concurrency safe
var Notices Notificater = &notification{}

func (uc *notification) Push(msg string) {
	if !Enable {
		return
	}
	uc.Hub <- msg
}

func (uc *notification) Listen(msgChan chan string) {
	uc.Hub = msgChan
	for {
		msg := <-msgChan
		err := uc.WriteMessage(websocket.TextMessage, []byte(msg))

		if err != nil {
			log.Fatal(err)
		}
	}
}

// Wshandler will upgrade http connection to websocket protocol
func Wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v", err)
		Enable = false
		return
	}

	Notices = &notification{conn, make(chan string)}
}
