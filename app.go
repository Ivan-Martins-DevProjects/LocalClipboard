package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	go a.StartServer()
	time.Sleep(2 * time.Second)

	url := "ws://localhost:5858"

	timeout := websocket.Dialer{
		HandshakeTimeout: 5 * time.Second,
	}

	conn, _, err := timeout.Dial(url, nil)
	if err != nil {
		log.Printf("Erro de conexão: %v", err)
	}
	if conn == nil {
		return
	}

	log.Printf("Conexão estabelecida")

	go func() {
		time.Sleep(2 * time.Second)
		a.ListenContent(conn)
	}()
}

func (a *App) ListenContent(conn *websocket.Conn) {
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Erro: %v", err)
			break
		}

		runtime.EventsEmit(a.ctx, "novo_clipboard", string(message))
	}
}

func (a *App) CopyToClipboard(text string) string {
	err := a.writeNativeClipboard(text)
	if err != nil {
		return "Erro: " + err.Error()
	}

	return "Copiado!" + text
}

func (a *App) StartServer() {
	http.HandleFunc("/", a.handleWebSocket)

	log.Println("Servidor rodando em :5858")
	if err := http.ListenAndServe(":5858", nil); err != nil {
		log.Fatal("Erro ao iniciar o servidor: ", err)
	}
}
