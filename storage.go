package Pigeon

import "github.com/gorilla/websocket"


type MessagesRepo struct {
    messages []*Message
}

type WebSocketsRepo struct {
    connections []*websocket.Conn
}

var MessagesRepository = MessagesRepo{}
var WebSocketsRepository = WebSocketsRepo{}

func (r *MessagesRepo) Add(m *Message) {
    if r.messages == nil {
        r.messages = []*Message{}
    }
    r.messages = append(r.messages, m)
}