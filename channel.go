package Pigeon

type Channel interface {
    Send(m *IMessage)
}