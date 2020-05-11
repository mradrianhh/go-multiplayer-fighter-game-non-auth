package main

import (
	"github.com/mradrianhh/go-multiplayer-fighter-game/server/pkg/network"
	"github.com/mradrianhh/go-multiplayer-fighter-game/server/pkg/program"
)

var response int

func main() {
	go network.Listen()

	program.Run()
}
