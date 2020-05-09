package game

import (
	"fmt"
	"os"

	"github.com/mradrianhh/go-multiplayer-fighter-game/client/navigation"
	"github.com/mradrianhh/go-multiplayer-fighter-game/pkg/vars"
)

var state vars.State

func init() {
	state = vars.HOME
}

// StartGame initializes the game.
func StartGame() {
	Loop()
}

// Loop is the gameloop.
func Loop() {
	for {
		switch state {
		case vars.MAIN:
			err := navigation.Screens["MAINMENU"].Show(&state)
			checkError(err)
		case vars.LOGIN:
			err := navigation.Screens["LOGIN"].Show(&state)
			checkError(err)
		case vars.REGISTER:
			err := navigation.Screens["REGISTER"].Show(&state)
			checkError(err)
		case vars.HOME:
			err := navigation.Screens["HOME"].Show(&state)
			checkError(err)
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
