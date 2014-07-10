package geyser

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func Run(fqdn string) {
	h := http.Header{}
	conn, _, err := websocket.DefaultDialer.Dial(fqdn, h)
	if err != nil {
		log.Println(err)
	}
	for {
		if err := conn.WriteMessage(websocket.TextMessage, pingMsg); err != nil {
			log.Println(err)
			break
		}
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
	}
	conn.Close()
}
