package client

import (
	"log"
	"time"
	"os"
	"os/signal"
	"net/url"
	"flag"
	"github.com/gorilla/websocket"
	"com.sinlov/Golang-Android-Remote-Control/conf"
	"fmt"
	"com.sinlov/Golang-Android-Remote-Control/event"
	"github.com/google/flatbuffers/go"
)

var cli_addr *string

func Input_cli() {
	flag.Parse()
	log.SetFlags(0)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	config := new(conf.Config)
	config.InitConfig("conf/config.conf")
	daemon := config.Read("ServerSet", "daemon")
	port := config.Read("ServerSet", "port")
	cli_addr = flag.String("addr", daemon + ":" + port, "http service address")
	u := url.URL{Scheme: "ws", Host: *cli_addr, Path: "/"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer c.Close()
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Post time string ", t.String())
			//err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			//if err != nil {
			//	log.Println("write String: ", err)
			//	return
			//}
			builder := flatbuffers.NewBuilder(0)
			menuBtn := builder.CreateString("3")
			//homeBtn := builder.CreateString("1")
			//time_str := builder.CreateString(t.String())
			event.KeyEventStart(builder)
			event.KeyEventAddKeyEvent(builder, menuBtn)
			end := event.KeyEventEnd(builder)
			builder.Finish(end)
			keyEvent := builder.FinishedBytes()
			fmt.Println(keyEvent)
			err_b := c.WriteMessage(websocket.BinaryMessage, keyEvent)
			if err_b != nil {
				log.Println("write binary err: ", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")
		// To cleanly close a connection, a client should send a close
		// frame and wait for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
				select {
				case <-done:
				case <-time.After(time.Second):
				}
			c.Close()
			return
		}
	}
}
