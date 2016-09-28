# Info//

- just in develop!


# Library

```sh
go get -u -v github.com/gorilla/websocket
go get -u -v github.com/google/flatbuffers/go
```

https://github.com/gorilla/websocket

# Debug

Set config file `conf/config.conf`

```conf
[ServerSet]

daemon=127.0.0.1    // for you daemon
port=18080          // you port number
```


change `Main.go`

```golang
func main() {
	start_sever()
	// or
	//start_comamd()
}
```

- Run

```sh
go run Main.go
```

# Connect

## Web

- Server Address will show in command
- Web Client URL will show in command


## dark websocket or other client
/connect ws://youdaemon:yourport


# Use cli

```sh
go run cli.go
```

# Use Video Player

- jsmpeg.js is from https://github.com/phoboslab/jsmpeg

- broadway js is from https://github.com/mbebenita/Broadway


See in URL
 
http://127.0.0.1:18080/h264/h264.html

http://127.0.0.1:18080/mpeg1/mpeg1.html

if you change [conf/config.conf] you must change daemon
