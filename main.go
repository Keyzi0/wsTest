package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
)

const poloniexURL = "api2.poloniex.com"

var addr = flag.String("addr", poloniexURL, "http service address")

func main() {
	store := NewStore()

	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: *addr, Path: "/realtime"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	first := true
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			r:= parseMsg(message)
			store.Set(r)
			//log.Printf("recv: %s,\n %v", message, r)
			if first {
				first = false
				c.WriteMessage(websocket.TextMessage, []byte(`{"command": "subscribe", "channel": 1002}`))
			}
		}
	}()


	fmt.Println("loop")
	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("interrupt")
			// Вывожу содержимое стора при завершении работы(просто так)
			for idx, s := range store.GetAll() {
				log.Printf(`%v: %v \n`, idx, s)
			}

			log.Printf(`store\n`)
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:

			case <-time.After(time.Second):
			}
			return
		}
	}
}

