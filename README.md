# Info//

- just in develop!

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

#dark websocket
/connect ws://127.0.0.1:18080/echo
