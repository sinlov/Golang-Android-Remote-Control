package wsVideoPlay

import (
	"fmt"
	"os"
	"log"
	"time"
	"io"
	"bytes"
	"encoding/binary"
	"golang.org/x/net/websocket"
	"net/http"
	"com.sinlov/Golang-Android-Remote-Control/conf"
)

func Start_h264() {
	config := new(conf.Config)
	config.InitConfig("conf/config.conf")
	daemon := config.Read("ServerSet", "daemon")
	port := config.Read("ServerSet", "port")
	http_Dir := config.Read("ServerSet", "http_Dir")
	http.Handle("/wsh264", websocket.Handler(wsH264))
	http.Handle("/wsmpeg", websocket.Handler(wsMpeg1))
	http.Handle("/", http.FileServer(http.Dir(http_Dir)))

	err := http.ListenAndServe(daemon + ":" + port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func wsH264(ws *websocket.Conn) {
	fmt.Printf("new socket\n")
	config := new(conf.Config)
	config.InitConfig("conf/config.conf")
	filePath := config.Read("FileSet", "file.testH264")
	fmt.Println("wsH264 path", filePath)
	//fi, err := os.Open("./test.h264")
	fi, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()

	msg := make([]byte, 1024 * 512)
	lenBytes := make([]byte, 4)
	for {
		time.Sleep(40 * time.Millisecond)

		lenNum, err := fi.Read(lenBytes)
		if (err != nil && err != io.EOF) || lenNum != 4 {
			log.Println(err)
			time.Sleep(1 * time.Second)
			break
		}

		b_buf := bytes.NewBuffer(lenBytes)
		var lenreal int32
		binary.Read(b_buf, binary.LittleEndian, &lenreal)

		n, err := fi.Read(msg[0:lenreal])
		if (err != nil && err != io.EOF) || n != int(lenreal) {
			log.Println(err)
			time.Sleep(1 * time.Second)
			break
		}

		err = websocket.Message.Send(ws, msg[:n])
		if err != nil {
			log.Println(err)
			break
		}
	}

	log.Println("send over h264 socket\n")
}

func wsMpeg1(ws *websocket.Conn) {
	fmt.Printf("new socket\n")

	buf := make([]byte, 10)
	buf[0] = 'j'
	buf[1] = 's'
	buf[2] = 'm'
	buf[3] = 'p'
	buf[4] = 0x01
	buf[5] = 0x40
	buf[6] = 0x0
	buf[7] = 0xf0
	websocket.Message.Send(ws, buf[:8])

	config := new(conf.Config)
	config.InitConfig("conf/config.conf")
	filePath := config.Read("FileSet", "file.testMpeg")
	fmt.Println("mpeg path", filePath)
	//fi, err := os.Open("./test.mpeg")
	fi, err := os.Open(filePath)
	//fi, err := os.Open("./Screenshot_2016-09-19-05-27-27-695.mp4")
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()

	msg := make([]byte, 1024 * 1)
	for {
		time.Sleep(40 * time.Millisecond)
		n, err := fi.Read(msg)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if 0 == n {
			time.Sleep(1 * time.Second)
			break
		}
		err = websocket.Message.Send(ws, msg[:n])
		if err != nil {
			log.Println(err)
			break
		}
	}
	fmt.Printf("send over mpeg1 socket\n")
}