package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"

	"github.com/mradrianhh/go-multiplayer-fighter-game/pkg/models"
)

func main() {
	player := models.NewPlayer("Adrianhh")

	service := "0.0.0.0:1200"

	conn, err := net.Dial("tcp", service)
	checkError(err)

	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)

	encoder.Encode(player)

	var newPlayer models.Player
	decoder.Decode(&newPlayer)

	fmt.Println(newPlayer.Exp)

	fmt.Scanln()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
