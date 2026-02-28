package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)

func (a *App) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	clients[conn] = true
	if err != nil {
		log.Printf("Erro ao estabelecer rede Socket: %v", err)
		return
	}
	defer func() {
		delete(clients, conn)
		conn.Close()
	}()

	log.Println("Cliente conectado via rede local")

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			break
		}

		for client := range clients {
			client.WriteMessage(websocket.TextMessage, p)
		}
	}
}
