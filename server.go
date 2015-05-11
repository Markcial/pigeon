package Pigeon

import (
    "github.com/gorilla/pat"
    "github.com/gorilla/websocket"
    "net/http"
)

type Hub interface {
    GetClients() []*websocket.Conn
}

type WsHub struct {
    Hub
    clients []*websocket.Conn
}

func (hub *WsHub) GetClients()[]*websocket.Conn {
    return hub.clients
}

func (hub *WsHub) AddClient(conn *websocket.Conn) {
    hub.clients = append(hub.clients, conn)
}

func Serve(router *pat.Router) {
    err := http.ListenAndServe(":8080", router)
    if err != nil {
        panic("Error: " + err.Error())
    }
}