package program

import (
	"fmt"
	"os"

	"github.com/mradrianhh/go-multiplayer-fighter-game/pkg/models"
	"github.com/mradrianhh/go-multiplayer-fighter-game/pkg/navigation"
	"github.com/mradrianhh/go-multiplayer-fighter-game/pkg/vars"
)

var state models.State

func init() {
	state.State = vars.Home
}

// Run initializes the game.
func Run() {
	Loop()
}

// Loop is the gameloop.
func Loop() {
	for {
		switch state.State {
		case vars.Main:
			err := navigation.RemoveScreensAndPush(&state, navigation.Screens["MAINMENU"])
			checkError(err)
		case vars.Login:
			err := navigation.RemoveScreensAndPush(&state, navigation.Screens["LOGIN"])
			checkError(err)
		case vars.Register:
			err := navigation.RemoveScreensAndPush(&state, navigation.Screens["REGISTER"])
			checkError(err)
		case vars.Home:
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
