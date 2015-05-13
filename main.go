package Pigeon

import (
    "fmt"
    "github.com/gorilla/websocket"
    "github.com/gorilla/pat"
    "net/http"
    "log"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

var wsHub = WsHub{}

type WSMessage struct {
  Name string `json:name`
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    //log.Printf("%v", conn)
    if err != nil {
        log.Println(err)
        return
    }

    wsHub.AddClient(conn)

    msg := WSMessage{}
    for {
        messageType, p, err := conn.ReadMessage()
        conn.ReadJSON(&msg)
        log.Println(err)
        log.Printf("%#v", msg)
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
  Serve(r)
}
