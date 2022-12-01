package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homepage")
}

func wsEndPoint(w http.ResponseWriter, r *http.Request) {
	upGrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrading connection to a websocket
	ws, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("error is endpoint", err)
	}
	// helpful log statement to show connection
	log.Println("Client Connected")
	reader(ws)
}

// reader will listenn for new messages
func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// pritn out message
		fmt.Println(string(p))
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}
func setRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndPoint)
}

func main() {
	fmt.Println("websocket")
	setRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
