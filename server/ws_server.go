package server

import (
	"flag"
	"net/http"
	"text/template"
	"log"
	"com.sinlov/Golang-Android-Remote-Control/conf"
	"fmt"
	"go/build"
	"path/filepath"
)

var (
	ws_addr *string
	assets = flag.String("assets", defaultAssetPath(), "path to assets")
	homeTempl *template.Template
)

func defaultAssetPath() string {
	p, err := build.Default.Import("gary.burd.info/go-websocket-chat", "", build.FindOnly)
	if err != nil {
		return "."
	}
	return p.Dir
}

func homeHandler(c http.ResponseWriter, req *http.Request) {
	homeTempl.Execute(c, req.Host)
}

func Start_ws_server() {
	config := new(conf.Config)
	config.InitConfig("conf/config.conf")
	daemon := config.Read("ServerSet", "daemon")
	port := config.Read("ServerSet", "port")
	fmt.Println("ws_server Init with: ", daemon, ":", port)
	ws_addr = flag.String("addr", daemon + ":" + port, "http service address")
	flag.Parse()
	homeTempl = template.Must(template.ParseFiles(filepath.Join(*assets, "server/home.html")))
	go h.run()
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ws", wsHandler)
	if err := http.ListenAndServe(*ws_addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}