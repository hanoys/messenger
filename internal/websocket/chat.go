package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("error while reading message: %v\n", err)
			return
		}

		log.Printf("message: %v\n", string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Printf("error while writing message: %v\n", err)
			return
		}
	}
}

func UpgradeConnection(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error on upgrade: %v\n", err)
		return
	}

	log.Println("client connected")
	reader(ws)
}
