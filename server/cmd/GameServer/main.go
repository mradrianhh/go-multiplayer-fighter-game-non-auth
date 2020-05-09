package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"

	"github.com/mradrianhh/go-multiplayer-fighter-game/pkg/models"
)

func main() {
	service := "0.0.0.0:1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		encoder := gob.NewEncoder(conn)
		decoder := gob.NewDecoder(conn)

		var player models.Player
		decoder.Decode(&player)

		player.UpdateExp(10)

		encoder.Encode(player)
		conn.Close() // we're finished
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
