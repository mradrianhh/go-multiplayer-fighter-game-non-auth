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
			err := navigation.RemoveScreensAndPush(&state, navigation.Screens["MAINMENU"])
			checkError(err)
		case vars.LOGIN:
			err := navigation.RemoveScreensAndPush(&state, navigation.Screens["LOGIN"])
			checkError(err)
		case vars.REGISTER:
			err := navigation.RemoveScreensAndPush(&state, navigation.Screens["REGISTER"])
			checkError(err)
		case vars.HOME:
			err := navigation.RemoveScreensAndPush(&state, navigation.Screens["HOME"])
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
