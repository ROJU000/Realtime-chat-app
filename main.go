// main.go
package main

import (
    "log"
    "net/http"

    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }
    defer conn.Close()

    for {
        // Read message from client
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }
        // Broadcast message to all connected clients
        for c := range clients {
            err := c.WriteMessage(messageType, p)
            if err != nil {
                log.Println(err)
                return
            }
        }
    }
}

func main() {
    http.HandleFunc("/ws", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

