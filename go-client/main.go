package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")
var mode = flag.String("mode", "ws", "Select the startup mode. The optional values are ws and http")

type Message struct {
	Wxid    string `json:"wxid"`
	Content string `json:"content"`
}

func wsClient() {
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", message)

		var msg Message
		json.Unmarshal(message, &msg)
		log.Printf("msg: %+v\n", msg)
	}
}

func httpServer() {
	g := gin.Default()
	g.POST("/callback", func(c *gin.Context) {
		var msg Message

		err := c.BindJSON(&msg)
		if err != nil {
			log.Printf("bind json faild: %s\n", err)
			return
		}

		log.Printf("msg: %+v\n", msg)
	})

	g.Run(*addr)
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	switch *mode {
	case "ws":
		wsClient()
	case "http":
		httpServer()
	}
}
