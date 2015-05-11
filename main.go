package Pigeon

import (
    "fmt"
    "github.com/gorilla/websocket"
    "github.com/gorilla/pat"
    "net/http"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

var wsHub = WsHub{}

func echoHandler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        //log.Println(err)
        return
    }
    wsHub.AddClient(conn)

    for {
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            return
        }

        err = conn.WriteMessage(messageType, p);
        if  err != nil {
            return
        }
    }
}

func composeMessageHandler(w http.ResponseWriter, r *http.Request) {

}

func viewMessageHandler(w http.ResponseWriter, r *http.Request) {

    fmt.Printf("%#v", &MessagesRepository)
    fmt.Printf("%#v", wsHub)
}

func addMessageHandler(w http.ResponseWriter, r *http.Request) {
    message := &Message{}
    MessagesRepository.Add(message)
}

var r = pat.New()

func Main() {
  r.Get("/echo", echoHandler)
  r.Get("/message/compose", composeMessageHandler)
  r.Get("/message/view", viewMessageHandler)
  r.Get("/message/add", addMessageHandler)
  r.Add("GET", "/", http.FileServer(http.Dir(".")))

  Serve(r)
}