package server

import (
	"flag"
	"net/http"
	"text/template"
	"log"
	"fmt"
	"path/filepath"
	"go/build"
	"github.com/sinlov/golang_utils/cfg"
)

var (
	ws_addr *string
	assets = flag.String("assets", defaultAssetPath(), "path to assets")
	homeTempl *template.Template
	daemon string
	port string
)

func defaultAssetPath() string {
	p, err := build.Default.Import("com.sinlov/Golang-Android-Remote-Control", "", build.FindOnly)
	if err != nil {
		return "."
	}
	return p.Dir
}

func homeHandler(w http.ResponseWriter, req *http.Request) {


	if req.URL.Path == "/" {
		http.Error(w, "404 Not found", 404)
		return
	}
	if req.Method != "GET" {
		http.Error(w, "405 Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	homeTempl.Execute(w, req.Host)
}

func Start_ws_server_by_router(routerSet string, htmlFilePath string) {
	if "" == daemon || "" == port {
		config := new(cfg.Cfg)
		config.InitCfg("conf/config.conf")
		daemon = config.Read("ServerSet", "daemon")
		port = config.Read("ServerSet", "port")
	}
	fmt.Println("WebSocket server Init with: ", daemon, ":", port)
	ws_addr = flag.String("addr", daemon + ":" + port, "http service address")
	flag.Parse()
	h := newHub()
	go h.run()
	//http.HandleFunc(routerSet, tempHandler)
	http.Handle(routerSet, wsHandler{h: h})
	http.Handle("/", http.FileServer(http.Dir(htmlFilePath)))
	if err := http.ListenAndServe(*ws_addr, nil); err != nil {
		log.Fatalf("ListenAndServe error %v , at ws_addr: %v", err, ws_addr)
	}
}

// Test info
// wsServerSet		use "/"
// jsClientRouterSet	use "/home"
// htmlFilePath		use "server/home.html"
func Start_ws_server_Test(wsServerSet string, jsClientRouterSet string, htmlFilePath string) {
	if "" == daemon || "" == port {
		config := new(cfg.Cfg)
		config.InitCfg("conf/config.conf")
		daemon = config.Read("ServerSet", "daemon")
		port = config.Read("ServerSet", "port")
	}
	fmt.Println("\n====\nWebSocket js client Init with\t\t http://" + daemon + ":" + port + jsClientRouterSet)
	ws_addr = flag.String("addr", daemon + ":" + port, "http service address")
	flag.Parse()
	homeTempl = template.Must(template.ParseFiles(filepath.Join(*assets, htmlFilePath)))
	h := newHub()
	go h.run()
	http.HandleFunc(jsClientRouterSet, homeHandler)
	http.Handle(wsServerSet, wsHandler{h: h})
	if err := http.ListenAndServe(*ws_addr, nil); err != nil {
		log.Fatalf("ListenAndServe error %v , at ws_addr: %v", err, ws_addr)
	}
}
