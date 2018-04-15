package socket

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait  = 2 * time.Second
	pongWait   = 10 * time.Second
	pingPeriod = (pongWait * 9) / 10
	maxMsgSize = 512
)

type notification struct {
	*websocket.Conn
	Hub chan string
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  512,
	WriteBufferSize: 1024,
}

// Notices globally holds user msg that we want to send, the Push func is concurrency safe
var Notices Notificater = new(notification)

func (uc *notification) Push(msg string) {
	if !Enable {
		return
	}
	uc.Hub <- msg
}

func (uc *notification) Listen(msgChan chan string) {
	uc.Hub = msgChan
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		close(msgChan)
	}()

	for {
		select {
		case msg := <-msgChan:
			uc.SetWriteDeadline(time.Now().Add(writeWait))
			w, err := uc.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Fatal(err)
				return
			}
			w.Write([]byte(msg))

			n := len(msgChan)
			for i := 0; i < n; i++ {
				w.Write([]byte(<-msgChan))
			}

			if err = w.Close(); err != nil {
				log.Fatal(err)
				return
			}

		case <-ticker.C:
			if err := uc.write(websocket.PingMessage, []byte{}); err != nil {
				log.Fatal(err)
				return
			}
		}
	}
}

func (uc *notification) ReadPump() {
	uc.SetReadLimit(maxMsgSize)
	uc.SetReadDeadline(time.Now().Add(pongWait))
	uc.SetPongHandler(func(string) error { uc.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, _, err := uc.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Fatal(err)
				break
			}
		}
	}
}

// write writes a message with the given message type and payload.
func (uc *notification) write(mt int, payload []byte) error {
	uc.SetWriteDeadline(time.Now().Add(writeWait))
	return uc.WriteMessage(mt, payload)
}

// Wshandler will upgrade http connection to websocket protocol
func Wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v", err)
		Enable = false
		return
	}
	defer conn.Close()

	n := new(notification)
	n.Conn = conn
	Notices = n
	go Notices.Listen(make(chan string, maxMsgSize))
	Notices.ReadPump()
}
