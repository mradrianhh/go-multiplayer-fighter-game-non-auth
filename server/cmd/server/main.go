package main

import (
	"github.com/mradrianhh/go-multiplayer-fighter-game/server/internal/pkg/network"
	"github.com/mradrianhh/go-multiplayer-fighter-game/server/internal/pkg/program"
)

var response int

func main() {
	network.Run()
	program.Run()
}
