package main

import (
	"flag"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")
var pathPrefix = flag.String("prefix", "/", "path prefix")

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	flag.Parse()

	log.SetFlags(0)

	if !strings.HasPrefix(*pathPrefix, "/") || !strings.HasSuffix(*pathPrefix, "/") {
		log.Fatalf("--prefix should start and finish with slash (/) or be equal it\n")
	}

	http.HandleFunc(*pathPrefix, homeHandler)
	http.HandleFunc((*pathPrefix)+"ws", wsHandler)

	log.Print("starting listen on ", *addr)
	log.Print("prefix path ", *pathPrefix)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	pageTpl, err := template.ParseFiles("./web/index.html")
	if err != nil {
		log.Println("ERROR: ", err.Error())
		return
	}

	data := map[string]string{}
	data["wsURL"] = (*pathPrefix) + "ws"

	err = pageTpl.Execute(w, data)
	if err != nil {
		log.Println("ERROR: ", err.Error())
		return
	}
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
