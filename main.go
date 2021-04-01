package main

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

func main() {
	flag.Parse()

	log.SetFlags(0)

	http.Handle("/", http.FileServer(http.Dir("./web/")))
	http.HandleFunc("/ws", wsHandler)

	log.Print("starting listen on ", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade: ", err)
		return
	}
	defer c.Close()
	for {
		message := time.Now().Format(time.RFC3339)

		err = c.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println("write: ", err)
			break
		}
		log.Println("write: ", message)

		rs := 500 + rand.Intn(1500)

		time.Sleep(time.Duration(rs) * time.Millisecond)
	}
}
