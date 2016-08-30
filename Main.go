package main

import (
	"com.sinlov/Golang-Android-Remote-Control/test"
	"flag"
	"log"
	"net/http"
	"os/exec"
	"com.sinlov/Golang-Android-Remote-Control/server"
)

func main() {
	server.Start_ws_server()
	//start_sever()
}

func start_sever() {
	flag.Parse()
	log.SetFlags(0)
	test.InitServer()
	http.HandleFunc("/", test.Echo)
	http.HandleFunc("/home", test.Home)
	log.Fatal(http.ListenAndServe(*test.Addr, nil))
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
	test.InitCommand(cmdPath)
	http.HandleFunc("/", test.ServeHome)
	http.HandleFunc("/ws", test.ServeWs)
	log.Fatal(http.ListenAndServe(*test.Command_addr, nil))
}
