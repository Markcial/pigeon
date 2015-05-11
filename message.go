package Pigeon

type IMessage interface {
    Send()
}

type Message struct {
    IMessage
    sender string
    receiver string
    channel Channel
    body string
}