package main

import (
	"com.sinlov/Golang-Android-Remote-Control/server"
	"flag"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	start_sever()
	// start_comamd()
}

func start_sever() {
	flag.Parse()
	log.SetFlags(0)
	server.InitServer()
	http.HandleFunc("/echo", server.Echo)
	http.HandleFunc("/", server.Home)
	log.Fatal(http.ListenAndServe(*server.Addr, nil))
}

func start_comamd() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Fatal("must specify at least one argument")
	}
	var err error
	cmdPath, err := exec.LookPath(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}
	server.InitCommand(cmdPath)
	http.HandleFunc("/", server.ServeHome)
	http.HandleFunc("/ws", server.ServeWs)
	log.Fatal(http.ListenAndServe(*server.Command_addr, nil))
}
