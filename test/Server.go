package test

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net/http"
	"github.com/sinlov/golang_utils/cfg"
)

var Addr *string

var server_upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	return true
}}

func InitServer() {
	config := new(cfg.Cfg)
	config.InitCfg("conf/config.conf")
	daemon := config.Read("ServerSet", "daemon")
	port := config.Read("ServerSet", "port")
	fmt.Println("You Server Init: ", daemon, ":", port)
	Addr = flag.String("addr", daemon + ":" + port, "http service address")
}

func Echo(w http.ResponseWriter, r *http.Request) {
	c, err := server_upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	ws_info := "ws://" + r.Host
	fmt.Println("webSocket info: ", ws_info)
	homeTemplate.Execute(w, ws_info)
}

var homeTemplate = template.Must(template.ParseFiles("test/Server.html"))
